/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/tstrijdhorst/tflow/cmd"
	"github.com/tstrijdhorst/tflow/services"
)

func main() {
	services.ConfigService{}.InitConfig()
	cmd.Execute()
}
