package main

import (
	"log"
	"net/http"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/websocket"
	services "github.com/billowdev/exclusive-go-hexa/internal/core/services/websocket"
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.ServeHTTP(w, r)
		if err != nil {
			log.Println("WebSocket connection error:", err)
			return
		}

		go func() {
			msgService := services.NewWebSocketMessageService(conn) // Pass the connection directly

			for {
				err := msgService.ProcessMessage()
				if err != nil {
					log.Println("Error processing message:", err)
					conn.Close() // Close the WebSocket connection on error
					break
				}
			}
		}()
	})

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
