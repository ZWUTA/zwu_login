package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	var (
		username string
		password string
		isLogout bool
	)

	flag.StringVar(&username, "u", "zc1", "Username")
	flag.StringVar(&password, "p", "1", "Password")
	flag.BoolVar(&isLogout, "L", false, "Logout")
	flag.Parse()

	if isLogout {
		logout()
	} else {
		login(username, password)
	}
}

func login(username string, password string) {
	log.Println("Login as", username, "...")
	client := &http.Client{}
	var data = strings.NewReader("DDDDD=" + username + "&upass=" + password + "&0MKKey=%B5%C7%C2%BC+Login")
	req, err := http.NewRequest("POST", "http://192.168.255.19/0.htm", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "192.168.255.19")
	req.Header.Set("Origin", "http://192.168.255.19")
	req.Header.Set("Referer", "http://192.168.255.19/0.htm")
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
	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	status := strings.Contains(buf.String(), "You have successfully logged into our system")

	if status {
		log.Println("Login successfully!")
	} else {
		log.Fatal("Login failed!")
	}
}

func logout() {
	log.Println("Logout...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://192.168.255.19/F.htm", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "192.168.255.19")
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
	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	status := strings.Contains(buf.String(), "Msg=14")
	if status {
		log.Println("Logout successfully!")
	} else {
		log.Fatal("Logout failed!")
	}
}
