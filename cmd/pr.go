package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tstrijdhorst/tflow/services"
)

// prCmd represents the merge command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a PR for the current branch",
	Long:  `Creates a PR for the current branch with a title equal to the issue id and summary and a body containing a link to the issue`,
	Run: func(cmd *cobra.Command, args []string) {
		createPR()
	},
}

func createPR() {
	jiraService := services.JiraService{
		Username: viper.GetString("jira.username"),
		Token:    viper.GetString("jira.token"),
		URL:      viper.GetString("jira.url"),
		Key:      viper.GetString("jira.key"),
	}

	branchName := services.GitService{}.GetCurrentBranchName()
	issueId, err := jiraService.ExtractIssueId(branchName)

	if err != nil {
		panic(err)
	}

	issueSummary := jiraService.GetSummaryForIssueId(issueId)
	prTitle := fmt.Sprintf("%v %v", issueId, issueSummary)
	prBody := jiraService.FormatIssueURL(issueId)

	services.GitService{}.PushCurrentBranch()

	url := services.GitHubService{}.CreatePullRequest(prTitle,prBody)
	fmt.Println(url)
}

func init() {
	rootCmd.AddCommand(prCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
