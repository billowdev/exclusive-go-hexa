package main

import (
	"log"
	"net/http"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/websocket"
	services "github.com/billowdev/exclusive-go-hexa/internal/core/services/websocket"
)

func main() {
	wsServer := websocket.NewWebSocketServer()
	chatService := services.NewChatService()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsServer.ServeHTTP(w, r)
		if err != nil {
			log.Println("WebSocket connection error:", err)
			return
		}

		chatService.RegisterClient(conn)
		go chatService.ProcessMessages(conn)
	})

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
