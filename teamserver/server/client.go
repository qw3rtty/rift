package server

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Conn       *websocket.Conn
	IP         string
	LastSeen   time.Time
	UserAgent  string
}

type ClientManager struct {
	sync.Mutex
	clients map[string]*Client
}

var ClManager = &ClientManager{
	clients: make(map[string]*Client),
}

func (m *ClientManager) Register(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client.ID] = client
	SaveClientToDB(client)
}

func (m *ClientManager) Unregister(id string) {
	m.Lock()
	defer m.Unlock()
	delete(m.clients, id)
	DeleteClientFromDB(id)
}

func (m *ClientManager) List() []*Client {
	m.Lock()
	defer m.Unlock()
	list := make([]*Client, 0, len(m.clients))
	for _, a := range m.clients {
		list = append(list, a)
	}
	return list
}
