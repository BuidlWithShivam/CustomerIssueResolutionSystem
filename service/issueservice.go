package service

import (
	"CustomerIssueResolutionSystem/entity"
	"CustomerIssueResolutionSystem/models"
	"CustomerIssueResolutionSystem/repository"
	"CustomerIssueResolutionSystem/util"
	"errors"
	"fmt"
)

type IssueService struct {
	issueRepository repository.IssueRepository
	agentService    *AgentService
}

func NewIssueService(issueRepository repository.IssueRepository,
	agentService *AgentService) *IssueService {
	return &IssueService{
		issueRepository: issueRepository,
		agentService:    agentService,
	}
}

func (service *IssueService) GetIssueById(id string) (*models.IssueResponse, error) {
	issue, err := service.issueRepository.GetIssueById(id)
	if err != nil {
		return nil, err
	}
	return &models.IssueResponse{
		IssueId:       issue.ExternalId,
		Type:          issue.IssueType,
		Status:        issue.Status,
		Subject:       issue.Subject,
		Description:   issue.Description,
		UserEmail:     issue.UserEmail,
		TransactionId: issue.TransactionId,
		AgentId:       issue.AgentId,
		Resolutions:   issue.Resolutions,
	}, nil
}

func (service *IssueService) CreateIssue(request *models.CreateIssueRequest) (*models.IssueResponse, error) {
	issue := &entity.Issue{
		ExternalId:    util.GenerateIssueID(),
		IssueType:     request.Type,
		Status:        models.CREATED,
		Subject:       request.Subject,
		Description:   request.Description,
		UserEmail:     request.UserEmail,
		TransactionId: request.TransactionId,
		Resolutions:   []string{},
	}
	_, err := service.issueRepository.CreateIssue(issue)
	if err != nil {
		return nil, err
	}
	return &models.IssueResponse{
		IssueId:       issue.ExternalId,
		Type:          issue.IssueType,
		Status:        issue.Status,
		Subject:       issue.Subject,
		Description:   issue.Description,
		UserEmail:     issue.UserEmail,
		TransactionId: issue.TransactionId,
		Resolutions:   issue.Resolutions,
	}, nil
}

func (service *IssueService) UpdateIssue(request *models.UpdateIssueRequest) (*models.IssueResponse, error) {
	issue, err := service.issueRepository.GetIssueById(request.IssueId)
	if err != nil {
		return nil, err
	}
	if request.Status != "" {
		issue.Status = request.Status
	}
	if request.Resolution != "" {
		issue.Resolutions = append(issue.Resolutions, request.Resolution)
	}
	issue, err = service.issueRepository.UpdateIssue(issue)
	if err != nil {
		return nil, err
	}
	return &models.IssueResponse{
		IssueId:       issue.ExternalId,
		Type:          issue.IssueType,
		Status:        issue.Status,
		Subject:       issue.Subject,
		Description:   issue.Description,
		UserEmail:     issue.UserEmail,
		TransactionId: issue.TransactionId,
		AgentId:       issue.AgentId,
		Resolutions:   issue.Resolutions,
	}, nil
}

func (service *IssueService) ResolveIssue(request *models.ResolveIssueRequest) (*models.IssueResponse, error) {
	issue, err := service.issueRepository.GetIssueById(request.IssueId)
	if err != nil {
		return nil, err
	}
	issue.Status = models.RESOLVED
	if request.Resolution != "" {
		issue.Resolutions = append(issue.Resolutions, request.Resolution)
	}
	issue, err = service.issueRepository.UpdateIssue(issue)
	if err != nil {
		return nil, err
	}
	_, err = service.agentService.UpdateAgent(&models.UpdateAgentRequest{
		AgentId:        issue.AgentId,
		Status:         models.AVAILABLE,
		CurrentIssueId: "",
	})
	if err != nil {
		return nil, err
	}
	err = service.agentService.UpdateAgentsWorkingHistory(issue.AgentId, issue.ExternalId)
	if err != nil {
		return nil, err
	}
	return &models.IssueResponse{
		IssueId:       issue.ExternalId,
		Type:          issue.IssueType,
		Status:        issue.Status,
		Subject:       issue.Subject,
		Description:   issue.Description,
		UserEmail:     issue.UserEmail,
		TransactionId: issue.TransactionId,
		AgentId:       issue.AgentId,
		Resolutions:   issue.Resolutions,
	}, nil
}

func (service *IssueService) GetIssues(filter models.IssueFilter) ([]models.IssueResponse, error) {
	issues, err := service.issueRepository.GetIssues(filter)
	if err != nil {
		return nil, err
	}
	issuesResponse := make([]models.IssueResponse, len(issues))
	for key, issue := range issues {
		issuesResponse[key] = models.IssueResponse{
			IssueId:       issue.ExternalId,
			Type:          issue.IssueType,
			Status:        issue.Status,
			Subject:       issue.Subject,
			Description:   issue.Description,
			UserEmail:     issue.UserEmail,
			TransactionId: issue.TransactionId,
			AgentId:       issue.AgentId,
			Resolutions:   issue.Resolutions,
		}
	}
	return issuesResponse, nil
}

func (service *IssueService) AssignIssue(issueId string) (string, error) {
	issue, err := service.issueRepository.GetIssueById(issueId)
	if err != nil {
		return "", err
	}
	if issue.Status == models.ASSIGNED || issue.Status == models.RESOLVED {
		return "", errors.New("issue already assigned")
	}
	agent, err := service.agentService.GetAvailableAgent(issue.IssueType)
	if err != nil {
		issue.Status = models.WAITING
		fmt.Println(err.Error())
		return "", nil
	}
	issue.AgentId = agent.AgentId
	issue.Status = models.ASSIGNED
	agent, err = service.agentService.UpdateAgent(&models.UpdateAgentRequest{
		AgentId:        agent.AgentId,
		Status:         models.WORKING,
		CurrentIssueId: issue.ExternalId,
	})
	if err != nil {
		return "", err
	}
	return agent.AgentId, nil
}
