/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tstrijdhorst/JFlow/cmd"
	"os"
)

func main() {
	initConfig()
	cmd.Execute()
}

const (
	FileName  = "config.yml"
	Directory = ".config/jflow"
)

func initConfig() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(fmt.Errorf("Could not find user homedir %w", err))
	}

	viper.SetConfigName(FileName)
	viper.AddConfigPath(homeDir + "/" + Directory)
	err = viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		}

		panic(fmt.Errorf("Fatal error in config file: %w \n", err))
	}
}
