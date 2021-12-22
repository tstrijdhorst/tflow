package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func main() {
	initConfig()
}

func initConfig() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error in config file: %w \n", err))
	}
}

func getJiraIssueSummary(id string) string {
	tp := jira.BasicAuthTransport{
		Username: viper.GetString("jira.username"),
		Password: viper.GetString("jira.token"),
	}

	client, err := jira.NewClient(tp.Client(), viper.GetString("jira.url"))

	//@todo how to recognize auth failure since apparently it doesnt return an error for that

	if err != nil {
		fmt.Printf("Error: %w", err)
	}

	issue, _, err := client.Issue.Get(id, nil)

	return issue.Fields.Summary
}
