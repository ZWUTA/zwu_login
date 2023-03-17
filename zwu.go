package main

import (
	"errors"
	"flag"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var client = http.DefaultClient

func main() {
	var (
		username      string
		password      string
		forceTryLogin bool
		logoutOpt     bool
		statusOpt     bool
	)

	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&password, "p", "", "Password")
	flag.BoolVar(&forceTryLogin, "f", false, "Force try login even though logged in")
	flag.BoolVar(&logoutOpt, "L", false, "Perform Logout operation")
	flag.BoolVar(&statusOpt, "S", false, "Perform GetStatus operation")
	flag.Parse()

	if logoutOpt {
		if err := logout(); err != nil {
			log.Println(err)
		} else {
			log.Println("success")
		}
	} else if statusOpt {
		if err := status(); err != nil {
			log.Println(err)
		}
	} else if username == "" || password == "" {
		log.Println("username and password can not be empty")
	} else {
		if _, err := getInfo(); err != nil || forceTryLogin {
			if err := login(username, password); err != nil {
				log.Println(err)
			} else {
				log.Println("success")
			}
			return
		}
		log.Println("skip: device already authed")
		log.Println(`if you really want to login again, retry with "-f" arg`)

	}
}

func login(username string, password string) error {
	log.Println("login as", username, "...")

	var data = strings.NewReader(`DDDDD=` +
		username +
		`&upass=` +
		password +
		`&0MKKey=%B5%C7%C2%BC+Login`)
	req, err := http.NewRequest("POST", "http://192.168.255.19/0.htm", data)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	newReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyText, err := io.ReadAll(newReader)
	if err != nil {
		return err
	}

	if strings.Contains(string(bodyText), "You have successfully logged into our system.") {
		return nil
	} else if strings.Contains(string(bodyText), "msga='userid error1'") {
		return errors.New("username " + username + " not exist")
	} else if strings.Contains(string(bodyText), "msga='userid error2'") {
		return errors.New("this account is banned")
	} else if strings.Contains(string(bodyText), "msga='ldap auth error'") {
		return errors.New("username or password error")
	} else if strings.Contains(string(bodyText), "msga='[02], 本帐号只能在指定 IP 段使用'") {
		return errors.New("this account only available in specified IP range")
	} else {
		return errors.New("unknown error")
	}
}

func logout() error {
	log.Println("logout...")

	req, err := http.NewRequest("GET", "http://192.168.255.19/F.htm", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	newReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyText, err := io.ReadAll(newReader)
	if err != nil {
		return err
	}

	if strings.Contains(string(bodyText), "Msg=14") {
		return nil
	} else {
		return errors.New("unknown error")
	}
}

func getInfo() (string, error) {
	req, err := http.NewRequest("GET", "http://192.168.255.19/", nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	newReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyText, err := io.ReadAll(newReader)
	if err != nil {
		return "", err
	}

	if strings.Contains(string(bodyText), `location.href="0.htm"`) {
		return string(bodyText), errors.New("device is not authed")
	} else {
		return string(bodyText), nil
	}
}

func status() error {
	bodyText, err := getInfo()
	if err != nil {
		return err
	}

	time := strings.Split(strings.Split(bodyText, "time='")[1], "';")[0]
	flow := strings.Split(strings.Split(bodyText, "flow='")[1], "';")[0]
	time = strings.Trim(time, " ")
	flow = strings.Trim(flow, " ")

	MB, _ := strconv.ParseFloat(flow, 64)
	MB /= 1024

	// less than 1h, print minutes, else if less than 24hours print hours, else print days
	if timeInt, err := strconv.Atoi(time); err == nil {
		if timeInt < 60 {
			log.Println("time:", time, "Min")
		} else if timeInt < 60*24 {
			log.Println("time:", timeInt/60, "Hour")
		} else {
			log.Println("time:", timeInt/60/24, "Day")
		}
	} else {
		log.Println("time:", time)
	}
	// less than 1GB, print MB, else if less than 1TB print GB, else print TB
	if MB < 1024 {
		log.Println("flow:", MB, "MB")
	} else if MB < 1024*1024 {
		log.Println("flow:", MB/1024, "GB")
	} else {
		log.Println("flow:", MB/1024/1024, "TB")
	}
	return nil
}
