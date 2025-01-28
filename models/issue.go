package models

type IssueType string

const (
	PaymentRelated    IssueType = "PaymentRelated"
	MutualFundRelated           = "MutualFundRelated"
	GoldRelated                 = "GoldRelated"
	InsuranceRelated            = "InsuranceRelated"
)

type IssueStatus string

const (
	RESOLVED IssueStatus = "RESOLVED"
	ASSIGNED IssueStatus = "ASSIGNED"
	WAITING  IssueStatus = "WAITING"
	CREATED  IssueStatus = "CREATED"
)

type CreateIssueRequest struct {
	TransactionId string    `json:"transactionId"`
	Type          IssueType `json:"type"`
	Subject       string    `json:"subject"`
	Description   string    `json:"description"`
	UserEmail     string    `json:"userEmail"`
}

type UpdateIssueRequest struct {
	IssueId    string      `json:"issueId"`
	Resolution string      `json:"resolution"`
	Status     IssueStatus `json:"status"`
}

type ResolveIssueRequest struct {
	IssueId    string `json:"issueId"`
	Resolution string `json:"resolution"`
}

type IssueFilter struct {
	Type      string `json:"type"`
	UserEmail string `json:"userEmail"`
	Status    string `json:"status"`
}

type IssueResponse struct {
	IssueId       string      `json:"issueId"`
	Type          IssueType   `json:"type"`
	Status        IssueStatus `json:"status"`
	Subject       string      `json:"subject"`
	Description   string      `json:"description"`
	UserEmail     string      `json:"userEmail"`
	TransactionId string      `json:"transactionId"`
	AgentId       string      `json:"agentId"`
	Resolutions   []string    `json:"resolutions"`
}
