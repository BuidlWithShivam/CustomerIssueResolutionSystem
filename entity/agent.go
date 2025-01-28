package entity

import "CustomerIssueResolutionSystem/models"

type Agent struct {
	User
	IssueTypes     []models.IssueType `json:"issueTypes"`
	Status         models.AgentStatus `json:"status"`
	CurrentIssueId string             `json:"currentIssueId"`
}

type AgentWorkingHistory struct {
	AgentId string   `json:"agentId"`
	Issues  []string `json:"issues"`
}
