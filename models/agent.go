package models

type AgentStatus string

const (
	AVAILABLE   AgentStatus = "AVAILABLE"
	UNAVAILABLE             = "UNAVAILABLE"
	WORKING                 = "WORKING"
)

type CreateAgentRequest struct {
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	IssueTypes []IssueType `json:"issueTypes"`
}

type UpdateAgentRequest struct {
	AgentId        string      `json:"agentId"`
	Status         AgentStatus `json:"status"`
	CurrentIssueId string      `json:"currentIssueId"`
}

type AgentResponse struct {
	AgentId    string      `json:"agentId"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	IssueTypes []IssueType `json:"issueTypes"`
}

type GetAgentRequest struct {
	AgentId   string      `json:"agentId"`
	Email     string      `json:"email"`
	IssueType IssueType   `json:"issueType"`
	Status    AgentStatus `json:"status"`
}

type GetAgentResponse struct {
	AgentId      string        `json:"agentId"`
	Email        string        `json:"email"`
	Status       AgentStatus   `json:"status"`
	CurrentIssue IssueResponse `json:"currentIssue"`
}

type AgentWorkingHistoryRequest struct {
	AgentId string `json:"agentId"`
}

type AgentWorkingHistoryResponse struct {
	AgentId string          `json:"agentId"`
	Issues  []IssueResponse `json:"issues"`
}
