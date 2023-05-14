package utils

import (
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"time"
)

type Config struct {
	Randomize       bool
	IntervalMinutes int
	ChangeOnHours   int
	ChangeOnGB      float64
}

type User struct {
	Enabled  bool
	Username string
	Password string
}

var runtimeViper = viper.New()

func init() {
	runtimeViper.SetConfigName("zwu")
	runtimeViper.SetConfigType("toml")
	runtimeViper.AddConfigPath(".")
}

func CreateConfig() error {
	runtimeViper.Set("config", Config{
		Randomize:       false,
		IntervalMinutes: 5,
		ChangeOnHours:   0,
		ChangeOnGB:      0.0,
	})
	runtimeViper.Set("user", []User{
		{
			Enabled:  false,
			Username: "username 1",
			Password: "password 1",
		}, {
			Enabled:  false,
			Username: "username 2",
			Password: "password 2",
		},
	})

	log.Println("Creating config file zwu.toml...")

	err := runtimeViper.SafeWriteConfig()
	if err != nil {
		return err
	}

	log.Println("Done!")
	return nil
}

func loadConfig() (Config, error) {
	log.Println("parsing config...")

	var config Config

	err := runtimeViper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = runtimeViper.UnmarshalKey("config", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func loadUsers() ([]User, error) {
	log.Println("parsing users...")
	err := runtimeViper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var users []User
	err = runtimeViper.UnmarshalKey("user", &users)
	if err != nil {
		return nil, err
	}

	var enabledUsers []User
	for _, user := range users {
		if user.Enabled {
			enabledUsers = append(enabledUsers, user)
		}
	}

	return enabledUsers, nil
}

func randomizeUserList(users []User) []User {
	log.Println("randomize user list...")

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	r.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	return users
}
