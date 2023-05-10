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

func CreateConfig() error {
	runtimeViper.SetConfigName("zwu")
	runtimeViper.SetConfigType("toml")
	runtimeViper.AddConfigPath(".")

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
