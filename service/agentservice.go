package service

import (
	"CustomerIssueResolutionSystem/entity"
	"CustomerIssueResolutionSystem/models"
	"CustomerIssueResolutionSystem/repository"
	"CustomerIssueResolutionSystem/util"
)

type AgentService struct {
	agentRepository repository.AgentRepository
	issueRepository repository.IssueRepository
}

func NewAgentService(agentRepository repository.AgentRepository,
	issueRepository repository.IssueRepository) *AgentService {
	return &AgentService{agentRepository: agentRepository,
		issueRepository: issueRepository}
}

func (service *AgentService) AddAgent(request *models.CreateAgentRequest) (*models.AgentResponse, error) {
	agent := new(entity.Agent)
	agent.ExternalId = util.GenerateAgentID()
	agent.Type = entity.AGENT
	agent.Email = request.Email
	agent.Name = request.Name
	agent.Status = models.AVAILABLE
	agent.IssueTypes = request.IssueTypes
	_, err := service.agentRepository.CreateAgent(agent)
	if err != nil {
		return nil, err
	}
	return &models.AgentResponse{
		AgentId:    agent.ExternalId,
		Name:       agent.Name,
		Email:      agent.Email,
		IssueTypes: agent.IssueTypes,
	}, nil
}

func (service *AgentService) ViewAgentsWorkingHistory(request *models.AgentWorkingHistoryRequest) (*models.AgentWorkingHistoryResponse, error) {
	agentWorkingHistory, err := service.agentRepository.GetAgentWorkingHistory(request.AgentId)
	if err != nil {
		return nil, err
	}
	var issues []models.IssueResponse
	for _, issueId := range agentWorkingHistory.Issues {
		issue, err := service.issueRepository.GetIssueById(issueId)
		if err != nil {
			return nil, err
		}
		issues = append(issues, models.IssueResponse{
			IssueId:       issue.ExternalId,
			Type:          issue.IssueType,
			Status:        issue.Status,
			Subject:       issue.Subject,
			Description:   issue.Description,
			UserEmail:     issue.UserEmail,
			TransactionId: issue.TransactionId,
			Resolutions:   issue.Resolutions,
		})
	}
	return &models.AgentWorkingHistoryResponse{
		AgentId: request.AgentId,
		Issues:  issues,
	}, nil
}

func (service *AgentService) UpdateAgentsWorkingHistory(agentId, issueId string) error {
	_, err := service.agentRepository.UpdateAgentWorkingHistory(agentId, issueId)
	if err != nil {
		return err
	}
	return nil
}

func (service *AgentService) GetAvailableAgent(issueType models.IssueType) (*models.AgentResponse, error) {
	agent, err := service.agentRepository.GetAvailableAgent(issueType)
	if err != nil {
		return nil, err
	}
	return &models.AgentResponse{
		AgentId:    agent.ExternalId,
		Name:       agent.Name,
		Email:      agent.Email,
		IssueTypes: agent.IssueTypes,
	}, nil
}

func (service *AgentService) UpdateAgent(request *models.UpdateAgentRequest) (*models.AgentResponse, error) {
	agent, err := service.agentRepository.FindAgentByID(request.AgentId)
	if err != nil {
		return nil, err
	}
	if agent.Status != "" {
		agent.Status = request.Status
	}
	if request.CurrentIssueId != "" {
		agent.CurrentIssueId = request.CurrentIssueId
	}
	agent, err = service.agentRepository.UpdateAgent(agent)
	if err != nil {
		return nil, err
	}
	return &models.AgentResponse{
		AgentId:    agent.ExternalId,
		Name:       agent.Name,
		Email:      agent.Email,
		IssueTypes: agent.IssueTypes,
	}, nil
}
