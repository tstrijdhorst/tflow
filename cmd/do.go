package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tstrijdhorst/tflow/services"
	"strings"
        "strconv"
)

const inProgressTransitionId string = "21"

var doCmd = &cobra.Command{
	Use:   "do <issueId> (with or without the KEY-)",
	Short: "Checkout a git branch for the given jira issueId. If it doesn't exist yet it is created.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doJiraIssueId(args[0])
	},
}

func doJiraIssueId(issueId string) {
	issueId = strings.TrimSpace(issueId)

        if _, err := strconv.Atoi(issueId); err == nil {
          issueId = viper.GetString("jira.key") + "-" + issueId;
        }

	jiraService := services.JiraService{
		Username: viper.GetString("jira.username"),
		Token:    viper.GetString("jira.token"),
		URL:      viper.GetString("jira.url"),
		Key:      viper.GetString("jira.key"),
	}

	jiraService.TransitionIssueId(issueId, inProgressTransitionId)
	
	issueSummary := jiraService.GetSummaryForIssueId(issueId)
	normalizedSummary := services.GitService{}.NormalizeForGitBranchName(issueSummary)

	branchName := issueId
	if normalizedSummary != "" {
		branchName += "/" + normalizedSummary
	}

	services.GitService{}.SwitchBranchCreateIfNotExists(branchName)
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
