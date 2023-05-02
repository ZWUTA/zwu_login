package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

func Logout(client *http.Client) error {
	log.Println("logout...")

	req, _ := http.NewRequest("GET", "http://192.168.255.19/F.htm", nil)
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

	if strings.Contains(bodyText, "Msg=14") {
		return nil
	} else {
		return errors.New("unknown error")
	}
}
