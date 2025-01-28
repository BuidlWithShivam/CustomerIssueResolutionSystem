package impl

import (
	"CustomerIssueResolutionSystem/entity"
	"CustomerIssueResolutionSystem/models"
	"CustomerIssueResolutionSystem/util"
	"errors"
)

type InMemoryAgentRepository struct {
	agents              map[string]*entity.Agent
	agentWorkingHistory map[string]*entity.AgentWorkingHistory
}

func NewInMemoryAgentRepository() *InMemoryAgentRepository {
	return &InMemoryAgentRepository{
		agents:              make(map[string]*entity.Agent),
		agentWorkingHistory: make(map[string]*entity.AgentWorkingHistory),
	}
}

func (i *InMemoryAgentRepository) CreateAgent(agent *entity.Agent) (*entity.Agent, error) {
	i.agents[agent.ExternalId] = agent
	i.agentWorkingHistory[agent.ExternalId] = &entity.AgentWorkingHistory{
		AgentId: agent.ExternalId,
		Issues:  []string{},
	}
	return agent, nil
}

func (i *InMemoryAgentRepository) FindAgentByID(id string) (*entity.Agent, error) {
	agent, ok := i.agents[id]
	if !ok {
		return nil, errors.New("Agent not found")
	}
	return agent, nil
}

func (i *InMemoryAgentRepository) UpdateAgent(agent *entity.Agent) (*entity.Agent, error) {
	_, ok := i.agents[agent.ExternalId]
	if !ok {
		return nil, errors.New("Agent  not found")
	}
	i.agents[agent.ExternalId] = agent
	return agent, nil
}

func (i *InMemoryAgentRepository) DeleteAgent(id string) error {
	_, ok := i.agents[id]
	if !ok {
		return errors.New("Agent  not found")
	}
	delete(i.agents, id)
	delete(i.agentWorkingHistory, id)
	return nil
}

func (i *InMemoryAgentRepository) GetAgentWorkingHistory(id string) (*entity.AgentWorkingHistory, error) {
	agentWorkingHistory, ok := i.agentWorkingHistory[id]
	if !ok {
		return nil, errors.New("AgentWorkingHistory not found")
	}
	return agentWorkingHistory, nil
}

func (i *InMemoryAgentRepository) UpdateAgentWorkingHistory(id string, issueId string) (*entity.AgentWorkingHistory, error) {
	agentWorkingHistory, ok := i.agentWorkingHistory[id]
	if !ok {
		return nil, errors.New("AgentWorkingHistory not found")
	}
	agentWorkingHistory.Issues = append(agentWorkingHistory.Issues, issueId)
	i.agentWorkingHistory[id] = agentWorkingHistory
	return agentWorkingHistory, nil
}

func (i *InMemoryAgentRepository) GetAvailableAgent(issueType models.IssueType) (*entity.Agent, error) {
	for _, agent := range i.agents {
		if agent.Status == models.AVAILABLE && util.IssueTypeContains(agent.IssueTypes, issueType) {
			return agent, nil
		}
	}
	return nil, errors.New("Agent not found")
}
