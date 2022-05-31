package services

import (
	"bytes"
	"fmt"
	"github.com/cli/safeexec"
	"log"
	"os/exec"
        "json"
)

type GitHubService struct {
}

func (g GitHubService) CreatePullRequest(title string) {
	ghBin, _ := safeexec.LookPath("gh")
	cmd := exec.Command(ghBin, "pr", "create", "--title", title, "--body", "")

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		//@todo nice error handling for common cases
		log.Fatal(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}

	fmt.Println(stdOut.String())
}

func (g GitHubService) MergePullRequest() {
	ghBin, _ := safeexec.LookPath("gh")
	cmd := exec.Command(ghBin, "pr", "merge", "--merge", "--auto")

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		//@todo nice error handling for common cases
		log.Fatal(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}

	fmt.Println(stdOut.String())
}


func (g GitHubService) GetBaseBranchName() string {
	ghBin, _ := safeexec.LookPath("gh")
	cmd := exec.Command(ghBin, "pr", "view", "--json", "baseRefName")

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		//@todo nice error handling for common cases
		log.Fatal(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}

        var js struct {
          baseRefName string `json:"baseRefName"`
        }

        err = json.Unmarshal(stdOut, &js)

	if err != nil {
		//@todo nice error handling for common cases
		log.Fatal(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}

        return js.baseRefName
}

