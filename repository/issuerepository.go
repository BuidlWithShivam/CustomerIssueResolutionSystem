package repository

import (
	"CustomerIssueResolutionSystem/entity"
	"CustomerIssueResolutionSystem/models"
)

type IssueRepository interface {
	CreateIssue(issue *entity.Issue) (*entity.Issue, error)
	UpdateIssue(issue *entity.Issue) (*entity.Issue, error)
	GetIssueById(id string) (*entity.Issue, error)
	GetIssues(filter models.IssueFilter) ([]entity.Issue, error)
}
