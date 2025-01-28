package repository

import (
	"CustomerIssueResolutionSystem/entity"
	"CustomerIssueResolutionSystem/models"
)

type AgentRepository interface {
	CreateAgent(agent *entity.Agent) (*entity.Agent, error)
	FindAgentByID(id string) (*entity.Agent, error)
	UpdateAgent(agent *entity.Agent) (*entity.Agent, error)
	DeleteAgent(id string) error
	GetAgentWorkingHistory(id string) (*entity.AgentWorkingHistory, error)
	UpdateAgentWorkingHistory(id string, issueId string) (*entity.AgentWorkingHistory, error)
	GetAvailableAgent(issueType models.IssueType) (*entity.Agent, error)
}
