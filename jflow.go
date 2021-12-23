package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/viper"
	"gopkg.in/andygrunwald/go-jira.v1"
	"os"
	"regexp"
	"strings"
)

func main() {
	initConfig()
	issueId := ""
	branchName := issueId
	normalizedSummary := normalizeForGitBranchName(getJiraIssueSummary(issueId))

	if normalizedSummary != "" {
		branchName += "/" + normalizedSummary
	}

	createNewGitBranch(branchName)
}

func normalizeForGitBranchName(s string) string {
	s = strings.ToLower(s)

	stripWhiteSpaceRegex := regexp.MustCompile(`\s`)
	s = stripWhiteSpaceRegex.ReplaceAllString(s, "_")

	stripIllegalCharsRegex := regexp.MustCompile(`[^a-z0-9.\-_/]+`)
	s = stripIllegalCharsRegex.ReplaceAllString(s, "")

	//Git branch names cannot start with '-' according to https://stackoverflow.com/a/3651867/298593
	return strings.TrimLeft(s, "-")
}

func initConfig() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error in config file: %w \n", err))
	}
}

func createNewGitBranch(name string) {
	workingDirectory, _ := os.Getwd()

	r, err := git.PlainOpen(workingDirectory)

	if err != nil {
		panic(fmt.Errorf("Fatal error in git repo: %w \n", err))
	}

	headRef, _ := r.Head()
	ref := plumbing.NewHashReference(plumbing.ReferenceName(fmt.Sprintf("refs/heads/%v", name)), headRef.Hash())

	// The created reference is saved in the storage.
	err = r.Storer.SetReference(ref)
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
