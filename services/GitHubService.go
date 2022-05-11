package services

import (
	"bytes"
	"fmt"
	"github.com/cli/safeexec"
	"log"
	"os/exec"
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
