package services

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type GitHubService struct {
}

func (g GitHubService) CreatePullRequest(title string) {
	cmd := exec.Command("gh", "pr", "create", "-t "+title, "-b ")

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
