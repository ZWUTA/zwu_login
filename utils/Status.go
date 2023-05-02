package utils

import (
	"log"
	"net/http"
)

func Status(client *http.Client) error {
	time, flow, err := GetAccount(client)
	if err != nil {
		return err
	}

	if time < 60 {
		log.Println("time:", time, "Min")
	} else if time < 60*24 {
		log.Println("time:", time/60, "Hour")
	} else {
		log.Println("time:", time/60/24, "Day")
	}

	flow /= 1024
	if flow < 1024 {
		log.Println("flow:", flow, "MB")
	} else if flow < 1024*1024 {
		log.Println("flow:", flow/1024, "GB")
	} else {
		log.Println("flow:", flow/1024/1024, "TB")
	}

	return nil
}
