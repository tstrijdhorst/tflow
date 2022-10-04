package services

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
	"strings"
	"regexp"
)

type JiraService struct {
	Token    string
	Username string
	URL      string
	Key      string
	client   *jira.Client
}

func (j JiraService) GetSummaryForIssueId(issueId string) string {
	issue, _, err := j.getClient().Issue.Get(issueId, nil)

	if err != nil {
		panic(fmt.Errorf("Error: %w \n", err))
	}

	return issue.Fields.Summary
}

func (j JiraService) FormatIssueId(issueNumber int) string {
        return fmt.Sprintf("%s-%d",j.Key,issueNumber);
}

func (j JiraService) FormatIssueURL(issueId string) string {
        jiraURL := strings.TrimSuffix(j.URL,"/")
        return fmt.Sprintf("%v/browse/%v",jiraURL,issueId);
}

func (j JiraService) TransitionIssueId(issueId, transitionId string) {
	_, err := j.getClient().Issue.DoTransition(issueId, transitionId)

	if err != nil {
		panic(fmt.Errorf("Error: %w \n", err))
	}
}

func (j JiraService) ExtractIssueId(s string) (string, error) {
	r := regexp.MustCompile("(" + j.Key + "-[0-9]+)")

	if !r.MatchString(s) {
		return "", fmt.Errorf("Could not find IssueId in string: %v", s)
	}

	return r.FindString(s), nil
}

func (j *JiraService) getClient() *jira.Client {
	if j.client == nil {
		j.initClient()
	}

	return j.client
}

func (j *JiraService) initClient() {
	tp := jira.BasicAuthTransport{
		Username: j.Username,
		Password: j.Token,
	}

	client, err := jira.NewClient(tp.Client(), j.URL)

	//@todo how to recognize auth failure since apparently it doesnt return an error for that

	if err != nil {
		fmt.Printf("Error: %w", err)
	}

	j.client = client
}
