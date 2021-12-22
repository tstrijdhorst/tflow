package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	fmt.Printf("%v", viper.Get("jira.api-key"))
}

func initConfig() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error in config file: %w \n", err))
	}
}
