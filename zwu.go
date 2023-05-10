package main

import (
	"flag"
	"log"
	"net/http"
	"zwu/utils"
)

var (
	username string
	password string

	optForceReLogin bool
	optLogout       bool
	optStatus       bool
	optCreateConfig bool
)

func init() {
	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&password, "p", "", "Password")

	flag.BoolVar(&optForceReLogin, "f", false, "Force re-login even though logged in")
	flag.BoolVar(&optLogout, "L", false, "Perform Logout operation")
	flag.BoolVar(&optStatus, "S", false, "Perform GetStatus operation")
	flag.BoolVar(&optCreateConfig, "C", false, "Create ./zwu.toml template")

	flag.Parse()
}

func main() {
	client := http.DefaultClient

	switch {
	case optLogout:
		err := utils.Logout(client)
		if err != nil {
			log.Println(err)
		}
		break
	case optStatus:
		err := utils.Status(client)
		if err != nil {
			log.Println(err)
		}
		break
	case optCreateConfig:
		err := utils.CreateConfig()
		if err != nil {
			log.Println(err)
		}
		break
	default:
		if username == "" || password == "" {
			log.Println("username and password can not be empty")
			return
		}

		_, _, err := utils.GetAccount(client)
		if err == nil && !optForceReLogin {
			log.Println("skip: device already authed")
			log.Println(`if you really want to login again, retry with "-f" arg`)
			return
		}

		err = utils.Login(client, username, password)
		if err != nil {
			log.Println(err)
		}
	}

}
