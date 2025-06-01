package server

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Agent struct {
	ID         string
	Conn       *websocket.Conn
	IP         string
	LastSeen   time.Time
	UserAgent  string
}

type AgentManager struct {
	sync.Mutex
	agents map[string]*Agent
}

var Manager = &AgentManager{
	agents: make(map[string]*Agent),
}

func (m *AgentManager) Register(agent *Agent) {
	m.Lock()
	defer m.Unlock()
	m.agents[agent.ID] = agent
	SaveAgentToDB(agent)
}

func (m *AgentManager) Unregister(id string) {
	m.Lock()
	defer m.Unlock()
	delete(m.agents, id)
	DeleteAgentFromDB(id)
}

func (m *AgentManager) List() []*Agent {
	m.Lock()
	defer m.Unlock()
	list := make([]*Agent, 0, len(m.agents))
	for _, a := range m.agents {
		list = append(list, a)
	}
	return list
}
