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

	optLogout       bool
	optStatus       bool
	optCreateConfig bool
	optDaemon       bool
)

func init() {
	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&password, "p", "", "Password")

	flag.BoolVar(&utils.ForceReLogin, "f", false, "Force re-login even though logged in")
	flag.BoolVar(&utils.Randomize, "r", false, "Random select username in config")

	flag.BoolVar(&optLogout, "L", false, "Perform Logout operation")
	flag.BoolVar(&optStatus, "S", false, "Perform GetStatus operation")
	flag.BoolVar(&optCreateConfig, "C", false, "Create ./zwu.toml template")
	flag.BoolVar(&optDaemon, "D", false, "Run as daemon")

	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	client := http.DefaultClient

	if optLogout {
		return utils.Logout(client)
	}

	if optStatus {
		return utils.Status(client)
	}

	if optCreateConfig {
		return utils.CreateConfig()
	}

	if optDaemon {
		return utils.Daemon(client)
	}

	if username == "" || password == "" {
		return utils.LoginFromConfig(client)
	}

	return utils.Login(client, username, password)
}
