package services

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

type JiraService struct {
	Token    string
	Username string
	URL      string
	client   *jira.Client
}

func (s JiraService) GetSummaryForIssueId(issueId string) string {
	issue, _, err := s.getClient().Issue.Get(issueId, nil)

	if err != nil {
		panic(fmt.Errorf("Error: %w \n", err))
	}

	return issue.Fields.Summary
}

func (s *JiraService) getClient() *jira.Client {
	if s.client == nil {
		s.initClient()
	}

	return s.client
}

func (s *JiraService) initClient() {
	tp := jira.BasicAuthTransport{
		Username: s.Username,
		Password: s.Token,
	}

	client, err := jira.NewClient(tp.Client(), s.URL)

	//@todo how to recognize auth failure since apparently it doesnt return an error for that

	if err != nil {
		fmt.Printf("Error: %w", err)
	}

	s.client = client
}
