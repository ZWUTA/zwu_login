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
	var err error
	client := http.DefaultClient

	switch {
	case optLogout:
		err = utils.Logout(client)
		if err != nil {
			log.Println(err)
		}
	case optStatus:
		err = utils.Status(client)
		if err != nil {
			log.Println(err)
		}
	case optCreateConfig:
		err = utils.CreateConfig()
		if err != nil {
			log.Println(err)
		}
	default:
		if username != "" && password != "" {

			_, _, err = utils.GetAccount(client)
			if err == nil && !optForceReLogin {
				log.Println("skip: device already authed")
				log.Println(`if you really want to login again, retry with "-f" arg`)
				return
			}

			err = utils.Login(client, username, password)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("login successful")
			return
		}

		log.Println("username and password not provided")
		log.Println("try parsing from config file...")

		users, err := utils.LoadUsers()
		if err != nil {
			log.Println(err)
			return
		}

		if len(users) == 0 {
			log.Println("no user found")
			return
		}
		log.Println("found", len(users), "user(s)")

		for _, user := range users {
			_, _, err = utils.GetAccount(client)
			if err == nil && !optForceReLogin {
				log.Println("skip: device already authed")
				log.Println(`if you really want to login again, retry with "-f" arg`)
				return
			}

			err = utils.Login(client, user.Username, user.Password)
			if err != nil {
				log.Println("login failed")
				log.Println(err)
				continue
			}
			_, _, err = utils.GetAccount(client)
			if err == nil {
				log.Println("login successful")
				return
			}
		}

		log.Println("all login attempts failed")
	}
}
