package util

import "CustomerIssueResolutionSystem/models"

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func IssueTypeContains(a []models.IssueType, x models.IssueType) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
