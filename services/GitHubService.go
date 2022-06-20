package services

import (
	"bytes"
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
