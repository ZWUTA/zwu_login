package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

func Login(client *http.Client, username string, password string) error {
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
