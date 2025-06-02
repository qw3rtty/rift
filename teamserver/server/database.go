package server

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "./rift.db")
	if err != nil {
		log.Fatal("Database error:", err)
	}

	createClientTable()
	createAgentTable()
}

// ------------------------------------
// -- Client related functions
// ------------------------------------
func createClientTable() {
	table := `
	CREATE TABLE IF NOT EXISTS clients (
		id TEXT PRIMARY KEY,
		ip TEXT,
		user_agent TEXT,
		last_seen TEXT
	);`
	_, err := db.Exec(table)
	if err != nil {
		log.Fatal("Error on creating 'clients' table:", err)
	}
}

func SaveClientToDB(client *Client) {
	stmt := `
	INSERT OR REPLACE INTO clients(id, ip, user_agent, last_seen)
	VALUES (?, ?, ?, ?);`
	_, err := db.Exec(stmt, client.ID, client.IP, client.UserAgent, client.LastSeen.Format(time.RFC3339))
	if err != nil {
		log.Println("Error on saving:", err)
	}
}

func DeleteClientFromDB(id string) {
	_, err := db.Exec(`DELETE FROM clients WHERE id = ?;`, id)
	if err != nil {
		log.Println("Error on deleting client:", err)
	}
}

func LoadClientsFromDB() []*Client {
	rows, err := db.Query(`SELECT id, ip, user_agent, last_seen FROM clients;`)
	if err != nil {
		log.Println("Error on loading clients:", err)
		return nil
	}
	defer rows.Close()

	var clients []*Client
	for rows.Next() {
		var c Client 
		var lastSeen string
		err := rows.Scan(&c.ID, &c.IP, &c.UserAgent, &lastSeen)
		if err != nil {
			continue
		}
		c.LastSeen, _ = time.Parse(time.RFC3339, lastSeen)
		clients = append(clients, &c)
	}
	return clients
}


// ------------------------------------
// -- Agent related functions
// ------------------------------------
func createAgentTable() {
	table := `
	CREATE TABLE IF NOT EXISTS agents (
		id TEXT PRIMARY KEY,
		ip TEXT,
		user_agent TEXT,
		last_seen TEXT
	);`
	_, err := db.Exec(table)
	if err != nil {
		log.Fatal("Error on creating 'agents' table:", err)
	}
}

func SaveAgentToDB(agent *Agent) {
	stmt := `
	INSERT OR REPLACE INTO agents (id, ip, user_agent, last_seen)
	VALUES (?, ?, ?, ?);`
	_, err := db.Exec(stmt, agent.ID, agent.IP, agent.UserAgent, agent.LastSeen.Format(time.RFC3339))
	if err != nil {
		log.Println("Error on saving:", err)
	}
}

func DeleteAgentFromDB(id string) {
	_, err := db.Exec(`DELETE FROM agents WHERE id = ?;`, id)
	if err != nil {
		log.Println("Error on deleting agent:", err)
	}
}

func LoadAgentsFromDB() []*Agent {
	rows, err := db.Query(`SELECT id, ip, user_agent, last_seen FROM agents;`)
	if err != nil {
		log.Println("Error on loading agents:", err)
		return nil
	}
	defer rows.Close()

	var agents []*Agent
	for rows.Next() {
		var a Agent
		var lastSeen string
		err := rows.Scan(&a.ID, &a.IP, &a.UserAgent, &lastSeen)
		if err != nil {
			continue
		}
		a.LastSeen, _ = time.Parse(time.RFC3339, lastSeen)
		agents = append(agents, &a)
	}
	return agents
}
