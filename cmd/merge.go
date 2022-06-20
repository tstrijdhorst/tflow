package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tstrijdhorst/tflow/services"
)

var doneFlag bool
var checkoutBaseFlag bool

const doneTransitionId string = "31"

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge the current PR into it's base",
	Long: `Merge the current PR into it's base`,
	Run: func(cmd *cobra.Command, args []string) {
          mergePR(doneFlag, checkoutBaseFlag);
	},
}

func mergePR(setIssueToDone bool, checkoutBase bool) {
  services.GitHubService{}.MergePullRequest()

  if setIssueToDone {
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
        jiraService.TransitionIssueId(issueId, doneTransitionId)	
  }

  if checkoutBase {
      baseBranch := services.GitHubService{}.GetBaseBranchName()
      services.GitService{}.SwitchBranchCreateIfNotExists(baseBranch)
  }
} 

func init() {
	rootCmd.AddCommand(mergeCmd)
        mergeCmd.Flags().BoolVarP(&doneFlag, "done", "d", false, "Set issue to done after merge")
        mergeCmd.Flags().BoolVarP(&checkoutBaseFlag, "checkout-base", "c", false, "Checks out the basebranch after merge")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
