package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"github.com/cli/safeexec"
)

type GitHubService struct {
}

func (g GitHubService) CreatePullRequest(title string) string {
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

        return strings.TrimSpace(stdOut.String())
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
          BaseRefName string `json:"baseRefName"`
        }

        err = json.Unmarshal(stdOut.Bytes(), &js)

	if err != nil {
		//@todo nice error handling for common cases
		log.Fatal(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}

        return js.BaseRefName
}

