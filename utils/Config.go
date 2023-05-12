package utils

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Randomize       bool
	ChangeOnMinutes int
	ChangeOnFlow    float64
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
		ChangeOnMinutes: 0,
		ChangeOnFlow:    0.0,
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

func LoadConfig() (Config, error) {
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

func LoadUsers() ([]User, error) {
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
