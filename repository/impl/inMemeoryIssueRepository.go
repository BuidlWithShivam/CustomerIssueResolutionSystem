package impl

import (
	"CustomerIssueResolutionSystem/entity"
	"CustomerIssueResolutionSystem/models"
	"errors"
)

type InMemoryIssueRepository struct {
	issues map[string]*entity.Issue
}

func NewInMemoryIssueRepository() *InMemoryIssueRepository {
	return &InMemoryIssueRepository{
		issues: make(map[string]*entity.Issue),
	}
}

func (i *InMemoryIssueRepository) CreateIssue(issue *entity.Issue) (*entity.Issue, error) {
	i.issues[issue.ExternalId] = issue
	return issue, nil
}

func (i *InMemoryIssueRepository) UpdateIssue(issue *entity.Issue) (*entity.Issue, error) {
	issue, ok := i.issues[issue.ExternalId]
	if !ok {
		return nil, errors.New("issue not found")
	}
	i.issues[issue.ExternalId] = issue
	return issue, nil
}

func (i *InMemoryIssueRepository) GetIssueById(id string) (*entity.Issue, error) {
	issue, ok := i.issues[id]
	if !ok {
		return nil, errors.New("issue not found")
	}
	return issue, nil
}

func (i *InMemoryIssueRepository) GetIssues(filter models.IssueFilter) ([]entity.Issue, error) {
	issues := make([]entity.Issue, 0)
	for _, issue := range i.issues {
		typeFilter := true
		emailFilter := true
		statusFilter := true
		if filter.Type != "" && string(issue.IssueType) != filter.Type {
			typeFilter = false
		}
		if filter.UserEmail != "" && issue.UserEmail != filter.UserEmail {
			emailFilter = false
		}
		if filter.Status != "" && string(issue.Status) != filter.Status {
			statusFilter = false
		}

		if typeFilter && emailFilter && statusFilter {
			issues = append(issues, *issue)
		}
	}
	return issues, nil
}
