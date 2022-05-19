package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tstrijdhorst/tflow/services"
)

var doneFlag bool

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge the current PR into it's base",
	Long: `Merge the current PR into it's base`,
	Run: func(cmd *cobra.Command, args []string) {
          mergePR(doneFlag);
	},
}

func mergePR(setIssueToDone bool) {
  services.GitHubService{}.MergePullRequest()

  if setIssueToDone {
    //transition jira issue to done status
  }
} 

func init() {
	rootCmd.AddCommand(mergeCmd)
        mergeCmd.Flags().BoolVarP(&doneFlag, "done", "d", false, "Set issue to done after merge")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
