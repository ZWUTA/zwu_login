package utils

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func Login(client *http.Client, username string, password string) error {
	if !ForceReLogin {
		if _, _, err := GetAccount(client); err == nil {
			log.Println("skip: device already authed")
			log.Println(`if you really want to login again, retry with "-f" arg`)
			return nil
		}
	}

	log.Println("login as", username, "...")

	var data = strings.NewReader(`DDDDD=` +
		username +
		`&upass=` +
		password +
		`&0MKKey=%B5%C7%C2%BC+Login`)
	req, _ := http.NewRequest("POST", "http://192.168.255.19/0.htm", data)
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
	bodyText := GbkToUtf8(resp.Body)

	switch {
	case strings.Contains(bodyText, "You have successfully logged into our system."):
		log.Println("login success")
		return nil
	case strings.Contains(bodyText, "msga='userid error1'"):
		return errors.New("username " + username + " not exist")
	case strings.Contains(bodyText, "msga='userid error2'"):
		return errors.New("this account is banned")
	case strings.Contains(bodyText, "msga='ldap auth error'"):
		return errors.New("username or password error")
	case strings.Contains(bodyText, "msga='[02], 本帐号只能在指定 IP 段使用'"):
		return errors.New("this account only available in specified IP range")
	default:
		return errors.New("unknown error")
	}
}

func LoginFromConfig(client *http.Client) error {
	log.Println("username and password not provided")
	log.Println("try parsing from config file...")

	users, err := LoadUsers()
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return errors.New("no user found")
	}

	log.Println("found", len(users), "user(s)")

	if Randomize {
		log.Println("randomize user list")

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		r.Shuffle(len(users), func(i, j int) {
			users[i], users[j] = users[j], users[i]
		})
	}

	for _, user := range users {
		if err := Login(client, user.Username, user.Password); err == nil {
			return nil
		}
		log.Println("login failed")
	}

	return errors.New("all login attempts failed")
}
