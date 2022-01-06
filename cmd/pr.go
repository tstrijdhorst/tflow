package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tstrijdhorst/JFlow/services"
)

// prCmd represents the merge command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a PR for the current branch with a title equal to the issue id and summary",
	Long:  `Creates a PR for the current branch with a title equal to the issue id and summary`,
	Run: func(cmd *cobra.Command, args []string) {
		createPR()
	},
}

func createPR() {
	jiraService := services.JiraService{
		Username: viper.GetString("jira.username"),
		Token:    viper.GetString("jira.token"),
		URL:      viper.GetString("jira.url"),
	}

	fmt.Println("Extracting issue-id from branch name")
	branchName := services.GitService{}.GetCurrentBranchName()
	issueId, err := jiraService.ExtractIssueId(branchName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Fetching issue summary from Jira")
	issueSummary := jiraService.GetSummaryForIssueId(issueId)
	prTitle := fmt.Sprintf("%v %v", issueId, issueSummary)
	fmt.Println("Pr Title Generated: ", prTitle)

	fmt.Println("Pushing current branch")
	services.GitService{}.PushCurrentBranch()

	fmt.Println("Creating PR at GitHub")
	services.GitHubService{}.CreatePullRequest(prTitle)
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
