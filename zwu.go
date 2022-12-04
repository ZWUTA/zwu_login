package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	idx := len(os.Args)
	if idx < 2 {
		log.Fatal("Missing operation args!")
	}
	switch os.Args[1] {
	case "login":
		if idx < 4 {
			log.Println("Missing Username or Password! Program will be logged as default user!")
			login("zc1", "1")
			break
		}
		log.Println("Logging as " + os.Args[2])
		login(os.Args[2], os.Args[3])
		break
	case "logout":
		log.Println("Logout")
		logout()
		break
	default:
		log.Println("Incorrect operation args!")
	}
}

func login(username string, password string) {
	client := &http.Client{}
	var data = strings.NewReader("DDDDD=" + username + "&upass=" + password + "&0MKKey=%B5%C7%C2%BC+Login")
	req, err := http.NewRequest("POST", "http://192.168.255.19/0.htm", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "192.168.255.19")
	req.Header.Set("Content-Length", "43")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "http://192.168.255.19")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.5195.102 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Referer", "http://192.168.255.19/0.htm")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "close")
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
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Logging status: %t\n", strings.Contains(string(bodyText), "You have successfully logged into our system."))
}

func logout() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://192.168.255.19/F.htm", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "192.168.255.19")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.5195.102 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "close")
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
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(bodyText), "Msg=14") {
		log.Println("Logout successfully!")
	} else {
		log.Println("Unknown error!")
	}
}
