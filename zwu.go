package main

import (
	"flag"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	var (
		username string
		password string
		isLogout bool
		isStatus bool
	)
	var client = http.DefaultClient

	flag.StringVar(&username, "u", "zc1", "Username")
	flag.StringVar(&password, "p", "1", "Password")
	flag.BoolVar(&isLogout, "L", false, "Logout")
	flag.BoolVar(&isStatus, "S", false, "Status")
	flag.Parse()

	if isLogout && isStatus {
		log.Fatalln("more than ONE operation")
	}
	if isLogout {
		logout(client)
		return
	}
	if isStatus {
		status(client)
		return
	}
	login(username, password, client)
}

func login(username string, password string, client *http.Client) {
	log.Println("login as", username, "...")

	var data = strings.NewReader(`DDDDD=` +
		username +
		`&upass=` +
		password +
		`&0MKKey=%B5%C7%C2%BC+Login`)
	req, err := http.NewRequest("POST", "http://192.168.255.19/0.htm", data)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	newReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyText, err := io.ReadAll(newReader)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(bodyText), "You have successfully logged into our system.") {
		log.Println("success")
	} else if strings.Contains(string(bodyText), "msga='userid error1'") {
		log.Fatalln("username", username, "not exist")
	} else if strings.Contains(string(bodyText), "msga='ldap auth error'") {
		log.Fatalln("username or password error")
	} else if strings.Contains(string(bodyText), "msga='[02], 本帐号只能在指定 IP 段使用'") {
		log.Fatalln("this account only available in specified IP range")
	} else {
		log.Fatalln("unknown error")
	}
}

func logout(client *http.Client) {
	log.Println("logout...")

	req, err := http.NewRequest("GET", "http://192.168.255.19/F.htm", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	newReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyText, err := io.ReadAll(newReader)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(bodyText), "Msg=14") {
		log.Println("success")
	} else {
		log.Fatalln("unknown error")
	}
}

func status(client *http.Client) {
	req, err := http.NewRequest("GET", "http://192.168.255.19/", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	newReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyText, err := io.ReadAll(newReader)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(bodyText), `location.href="0.htm"`) {
		log.Fatalln("device is not authed")
	}
	time := strings.Split(strings.Split(string(bodyText), "time='")[1], "';")[0]
	flow := strings.Split(strings.Split(string(bodyText), "flow='")[1], "';")[0]
	time = strings.Trim(time, " ")
	flow = strings.Trim(flow, " ")

	MB, _ := strconv.ParseFloat(flow, 64)
	MB /= 1024
	log.Println("time:", time, "Min")
	log.Println("flow:", MB, "MB")
}
