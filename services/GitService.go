package services

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
	"regexp"
	"strings"
)

type GitService struct {
	repository *git.Repository
}

func (g GitService) GetCurrentBranchName() string {
	headRef, err := g.getRepository().Head()

	if err != nil {
		panic(fmt.Errorf("Unable to get HEAD: %w", err))
	}

	//@todo what happens when headref is a commit and not a branch?
	return headRef.Name().String()
}

//@todo bug, for some reason this deletes files that are locally ignored?
func (g GitService) CreateBranchIfNotExistsAndCheckout(name string) {
	worktree, err := g.getRepository().Worktree()

	if err != nil {
		panic(fmt.Errorf("Error getting worktree: %w", err))
	}

	b := plumbing.NewBranchReferenceName(name)

	// First try to checkout branch
	err = worktree.Checkout(&git.CheckoutOptions{Create: false, Force: false, Branch: b})

	if err != nil {
		// got an error - try to create it
		_ = worktree.Checkout(&git.CheckoutOptions{Create: true, Force: false, Branch: b})
	}
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

func (g *GitService) getRepository() *git.Repository {
	if g.repository == nil {
		g.initRepository()
	}

	return g.repository
}

func (g *GitService) initRepository() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(fmt.Errorf("Unable to get working directory %w", err))
	}

	r, err := git.PlainOpen(cwd)

	if err != nil {
		panic(fmt.Errorf("Fatal error in git repo: %w \n", err))
	}

	g.repository = r
}
