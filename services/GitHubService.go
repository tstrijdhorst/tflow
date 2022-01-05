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
	cmd := exec.Command("gh", "pr", "create", "-t '"+title+"'", "-b ''")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %q\n", out.String())
}
