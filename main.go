/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tstrijdhorst/JFlow/cmd"
)

func main() {
	initConfig()
	cmd.Execute()
}

func initConfig() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error in config file: %w \n", err))
	}
}
