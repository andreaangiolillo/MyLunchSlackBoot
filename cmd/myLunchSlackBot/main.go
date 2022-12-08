package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func newViper() error {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	return nil
}

func main() {
	if err := newViper; err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	token := viper.GetString("slack.token")
	if token == "" {
		panic("Missing slack token")
	}

	api := slack.New(token)
	user, err := api.GetUserInfo("U04ETL260C9")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
