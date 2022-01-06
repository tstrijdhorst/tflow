package services

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type GitService struct {
}

func (g GitService) PushCurrentBranch() {
	cmd := exec.Command("git", "push")

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		//@todo nice error handling for common cases
		log.Panic(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}
}

// GetCurrentBranchName Requires git >= 2.22
func (g GitService) GetCurrentBranchName() string {
	cmd := exec.Command("git", "branch", "--show-current")

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		log.Panic(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}

	return stdOut.String()
}

func (g GitService) SwitchBranchCreateIfNotExists(name string) {
	g.switchBranch(name, !g.branchExists(name))
}

func (g GitService) NormalizeForGitBranchName(s string) string {
	s = strings.ToLower(s)

	stripWhiteSpaceRegex := regexp.MustCompile(`\s`)
	s = stripWhiteSpaceRegex.ReplaceAllString(s, "_")

	stripIllegalCharsRegex := regexp.MustCompile(`[^a-z0-9.\-_/]+`)
	s = stripIllegalCharsRegex.ReplaceAllString(s, "")

	//Git branch names cannot start with '-' according to https://stackoverflow.com/a/3651867/298593
	return strings.TrimLeft(s, "-")
}

// requires git >= 2.23
func (g GitService) switchBranch(name string, create bool) {
	var cmd *exec.Cmd
	if create {
		cmd = exec.Command("git", "switch", "-c", name)
	} else {
		cmd = exec.Command("git", "switch", name)
	}

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		log.Panic(fmt.Errorf("ERROR: %v StdOut: %v StdErr: %v", err, stdOut.String(), stdErr.String()))
	}
}

func (g GitService) branchExists(name string) bool {
	err := exec.Command("git", "rev-parse", name).Run()
	return err == nil
}
