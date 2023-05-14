package utils

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Daemon(client *http.Client) error {
	log.Println("daemon mode")

	ForceReLogin = true
	config, err := loadConfig()
	users, err := loadUsers()
	if err != nil {
		return err
	}

	if !(config.IntervalMinutes > 0) {
		return errors.New("IntervalMinutes must be greater than 0")
	}

	log.Println("###################")
	log.Println("IntervalMinutes:", config.IntervalMinutes)

	if config.ChangeOnHours <= 0 {
		log.Println("ChangeOnHours: disabled")
	} else {
		log.Println("ChangeOnHours:", config.ChangeOnHours)
	}

	if config.ChangeOnGB <= 0 {
		log.Println("ChangeOnGB: disabled")
	} else {
		log.Println("ChangeOnGB:", config.ChangeOnGB)
	}

	log.Println("###################")
	log.Println("daemon mode started")

	if config.Randomize {
		users = randomizeUserList(users)
	}

	err = LoginUserList(client, users)
	if err != nil {
		log.Println(err)
	}

	ticker := time.NewTicker(time.Duration(config.IntervalMinutes) * time.Minute)
	for range ticker.C {

		minutes, flow, err := GetAccount(client)
		if err != nil {
			err := LoginUserList(client, users)
			if err != nil {
				log.Println(err)
			}
			continue
		}

		log.Println("hours:", minutes/60)
		log.Println("flow:", flow/1024/1024/1024, "GB")

		if minutes >= config.ChangeOnHours*60 && config.ChangeOnHours > 0 {
			log.Println("change user for time limit reached!")

			if config.Randomize {
				users = randomizeUserList(users)
			}

			err := LoginUserList(client, users)
			if err != nil {
				log.Println(err)
			}
		} else if flow/1024/1024/1024 >= config.ChangeOnGB && config.ChangeOnGB > 0 {
			log.Println("change user for flow limit reached!")

			if config.Randomize {
				log.Println("randomize user list...")

				source := rand.NewSource(time.Now().UnixNano())
				r := rand.New(source)
				r.Shuffle(len(users), func(i, j int) {
					users[i], users[j] = users[j], users[i]
				})
			}

			err := LoginUserList(client, users)
			if err != nil {
				log.Println(err)
			}
		}

	}

	return nil
}
