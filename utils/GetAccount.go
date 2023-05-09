package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetAccount(client *http.Client) (int, float64, error) {
	req, _ := http.NewRequest("GET", "http://192.168.255.19/", nil)
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	bodyText := GbkToUtf8(resp.Body)

	if strings.Contains(bodyText, `location.href="0.htm"`) {
		return 0, 0, errors.New("device is not authed")
	}

	time := strings.Split(strings.Split(bodyText, "time='")[1], "';")[0]
	flow := strings.Split(strings.Split(bodyText, "flow='")[1], "';")[0]
	time = strings.Trim(time, " ")
	flow = strings.Trim(flow, " ")

	timeInt, _ := strconv.Atoi(time)
	flowFloat64, _ := strconv.ParseFloat(flow, 64)

	return timeInt, flowFloat64, nil
}
