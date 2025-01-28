package util

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateIssueID() string {
	return "IS" + GenerateUUID()
}

func GenerateAgentID() string {
	return "AG" + GenerateUUID()
}
