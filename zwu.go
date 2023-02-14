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

	flag.StringVar(&username, "u", "zc1", "Username")
	flag.StringVar(&password, "p", "1", "Password")
	flag.BoolVar(&isLogout, "L", false, "Logout")
	flag.BoolVar(&isStatus, "S", false, "Status")
	flag.Parse()

	if isLogout && isStatus {
		log.Fatalln("more than ONE operation")
	}
	if isLogout {
		logout()
		return
	}
	if isStatus {
		status()
		return
	}
	login(username, password)
}

func login(username string, password string) {
	log.Println("login as", username, "...")

	client := &http.Client{}
	var data = strings.NewReader(`DDDDD=` +
		username +
		`&upass=` +
		password +
		`&0MKKey=%B5%C7%C2%BC+Login`)
	req, err := http.NewRequest("POST", "http://192.168.255.19/0.htm", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Origin", "http://192.168.255.19")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Referer", "http://192.168.255.19/0.htm")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
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
	} else {
		log.Fatalln("unknown error")
	}
}

func logout() {
	log.Println("logout...")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://192.168.255.19/F.htm", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Referer", "http://192.168.255.19/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
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

func status() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://192.168.255.19/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "myusername=2020015008; username=2020015008; smartdot=nnzzYSH123")
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
