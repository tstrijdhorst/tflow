package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tstrijdhorst/tflow/services"
)

var branchCmd = &cobra.Command{
	Use:   "branch <issueId>",
	Short: "Create and checkout a git branch for the given jira issueId",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createBranchFromJiraIssueId(args[0])
	},
}

func createBranchFromJiraIssueId(issueId string) {
	issueSummary := services.JiraService{
		Username: viper.GetString("jira.username"),
		Token:    viper.GetString("jira.token"),
		URL:      viper.GetString("jira.url"),
	}.GetSummaryForIssueId(issueId)
	normalizedSummary := services.GitService{}.NormalizeForGitBranchName(issueSummary)

	branchName := issueId
	if normalizedSummary != "" {
		branchName += "/" + normalizedSummary
	}

	services.GitService{}.SwitchBranchCreateIfNotExists(branchName)
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
