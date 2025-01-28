package entity

import "CustomerIssueResolutionSystem/models"

type Issue struct {
	Id            int                `json:"id"`
	ExternalId    string             `json:"externalId"`
	IssueType     models.IssueType   `json:"issueType"`
	Status        models.IssueStatus `json:"status"`
	Subject       string             `json:"subject"`
	Description   string             `json:"description"`
	UserEmail     string             `json:"userEmail"` // Purpose of this problem else user
	AgentId       string             `json:"agentId"`
	TransactionId string             `json:"transactionId"`
	Resolutions   []string           `json:"resolutions"`
}
