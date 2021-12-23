package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/andygrunwald/go-jira.v1"
	"os"
	"regexp"
	"strings"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch <issueId>",
	Short: "Create and checkout a git branch for the given jira issueId",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branch(args[0])
	},
}

func branch(issueId string) {
	branchName := issueId
	normalizedSummary := normalizeForGitBranchName(getJiraIssueSummary(issueId))

	if normalizedSummary != "" {
		branchName += "/" + normalizedSummary
	}

	createBranchIfNotExistsAndCheckout(branchName)
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

func createBranchIfNotExistsAndCheckout(name string) {
	workingDirectory, _ := os.Getwd()

	r, err := git.PlainOpen(workingDirectory)

	if err != nil {
		panic(fmt.Errorf("Fatal error in git repo: %w \n", err))
	}

	worktree, _ := r.Worktree()

	b := plumbing.NewBranchReferenceName(name)

	// First try to checkout branch
	err = worktree.Checkout(&git.CheckoutOptions{Create: false, Force: false, Branch: b})

	if err != nil {
		// got an error - try to create it
		_ = worktree.Checkout(&git.CheckoutOptions{Create: true, Force: false, Branch: b})
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

func init() {
	rootCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
