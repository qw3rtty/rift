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
	db, err = sql.Open("sqlite", "./agents.db")
	if err != nil {
		log.Fatal("Database error:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS agents (
		id TEXT PRIMARY KEY,
		ip TEXT,
		user_agent TEXT,
		last_seen TEXT
	);`
	_, err = db.Exec(createTable)
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
