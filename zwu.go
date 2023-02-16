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
		username string
		password string
	)

	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&password, "p", "", "Password")
	logoutOpt := flag.Bool("L", false, "Perform Logout operation")
	statusOpt := flag.Bool("S", false, "Perform GetStatus operation")
	flag.Parse()

	if *logoutOpt {
		err := logout()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("success")
		}
	} else if *statusOpt {
		err := status()
		if err != nil {
			log.Println(err)
		}
	} else if username == "" || password == "" {
		log.Println("username and password can not be empty")
	} else {
		err := login(username, password)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("success")
		}
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
	log.Println("time:", time, "Min")
	log.Println("flow:", MB, "MB")
	return nil
}
