package services

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"os"
)

const (
	FileName  = "config"
	Directory = ".config/tflow"
)

type ConfigService struct {
}

type configAnswers struct {
	URL      string
	Username string
	Key      string
	Token    string
}

func (c ConfigService) InitConfig() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(fmt.Errorf("Could not find user homedir %w", err))
	}

	viper.SetConfigName(FileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(homeDir + "/" + Directory)
	err = viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Lets generate one")
			answers := c.getConfigInteractive()

			viper.Set("jira.url", answers.URL)
			viper.Set("jira.username", answers.Username)
			viper.Set("jira.key", answers.Key)
			viper.Set("jira.token", answers.Token)

			configDirPath := homeDir + "/" + Directory + "/"
			fmt.Println("LOL")
			if _, err := os.Stat(configDirPath); os.IsNotExist(err) {
				err := os.Mkdir(configDirPath, 0700)
				if err != nil {
					panic(err)
				}
			}

			err := viper.WriteConfigAs(configDirPath + FileName + ".yml")
			if err != nil {
				panic(err)
			}
			return
		}

		panic(fmt.Errorf("Fatal error in config file: %w \n", err))
	}
}

func (c ConfigService) getConfigInteractive() configAnswers {
	//@todo validate these values and validate that auth works with jira
	// the questions to ask
	var qs = []*survey.Question{
		{
			Name:     "url",
			Prompt:   &survey.Input{Message: "What is the url of your jira cloud instance?"},
			Validate: survey.Required,
		},
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: "What username do you use?"},
			Validate: survey.Required,
		},
		{
			Name:     "key",
			Prompt:   &survey.Input{Message: "What key is your project using (e.g TFLOW)?"},
			Validate: survey.Required,
		},
		{
			Name:     "token",
			Prompt:   &survey.Input{Message: "What is your API token? (Generate one if necessary at: https://id.atlassian.com/manage-profile/security/api-tokens)"},
			Validate: survey.Required,
		},
	}

	var answers configAnswers
	err := survey.Ask(qs, &answers)

	if err != nil {
		panic(err)
	}

	return answers
}
