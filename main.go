package main

import (
	"CustomerIssueResolutionSystem/models"
	"CustomerIssueResolutionSystem/repository/impl"
	"CustomerIssueResolutionSystem/service"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	agentRepository := impl.NewInMemoryAgentRepository()
	issueRepository := impl.NewInMemoryIssueRepository()

	agentService := service.NewAgentService(agentRepository, issueRepository)
	issueService := service.NewIssueService(issueRepository, agentService)

	issue1, err := issueService.CreateIssue(&models.CreateIssueRequest{
		TransactionId: "T1",
		Type:          models.PaymentRelated,
		Subject:       "Payment Failed",
		Description:   "My payment failed but money is debited",
		UserEmail:     "testUser1@test.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue created : ", issue1)

	issue2, err := issueService.CreateIssue(&models.CreateIssueRequest{
		TransactionId: "T2",
		Type:          models.MutualFundRelated,
		Subject:       "Purchase Failed",
		Description:   "Unable to purchase Mutual Fund",
		UserEmail:     "testUser2@test.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue created : ", issue2)

	issue3, err := issueService.CreateIssue(&models.CreateIssueRequest{
		TransactionId: "T3",
		Type:          models.PaymentRelated,
		Subject:       "Payment Failed",
		Description:   "My payment failed but money is debited",
		UserEmail:     "testUser2@test.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue created : ", issue3)

	agent1, err := agentService.AddAgent(&models.CreateAgentRequest{
		Name:       "agent1@test.com",
		Email:      "Agent 1",
		IssueTypes: []models.IssueType{models.PaymentRelated},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Agent added : ", agent1)

	agent2, err := agentService.AddAgent(&models.CreateAgentRequest{
		Name:       "agent2@test.com",
		Email:      "Agent 2",
		IssueTypes: []models.IssueType{models.MutualFundRelated},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Agent added : ", agent2)

	agentId1, err := issueService.AssignIssue(issue1.IssueId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue assigned to agent : ", issue1.IssueId, agentId1)

	agentId2, err := issueService.AssignIssue(issue2.IssueId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue assigned to agent : ", issue2.IssueId, agentId2)

	agentId3, err := issueService.AssignIssue(issue3.IssueId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue assigned to agent : ", issue3.IssueId, agentId3)

	_, err = issueService.ResolveIssue(&models.ResolveIssueRequest{
		IssueId:    issue1.IssueId,
		Resolution: "Resolved by doing it",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue resolved : ", issue1.IssueId)

	agentId3, err = issueService.AssignIssue(issue3.IssueId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue assigned to agent : ", issue3.IssueId, agentId3)

	_, err = issueService.ResolveIssue(&models.ResolveIssueRequest{
		IssueId:    issue2.IssueId,
		Resolution: "Resolved by doing it 2",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue resolved : ", issue2.IssueId)

	_, err = issueService.ResolveIssue(&models.ResolveIssueRequest{
		IssueId:    issue3.IssueId,
		Resolution: "Resolved by doing it 3",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Issue resolved : ", issue2.IssueId)

	response, err :=
		agentService.ViewAgentsWorkingHistory(&models.AgentWorkingHistoryRequest{AgentId: agent1.AgentId})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	response, err =
		agentService.ViewAgentsWorkingHistory(&models.AgentWorkingHistoryRequest{AgentId: agent2.AgentId})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
