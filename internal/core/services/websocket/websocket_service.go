package services

import (
	"sync"

	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/websocket"
	"github.com/gorilla/websocket"
)

type ChatService struct {
    clients map[*websocket.Conn]bool
    mu      sync.Mutex
}

func NewChatService() *ChatService {
    return &ChatService{
        clients: make(map[*websocket.Conn]bool),
    }
}

func (s *ChatService) RegisterClient(conn *websocket.Conn) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.clients[conn] = true
}

func (s *ChatService) UnregisterClient(conn *websocket.Conn) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.clients, conn)
}

func (s *ChatService) BroadcastMessage(msg domain.WebSocketExampleMessage) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    for client := range s.clients {
        err := client.WriteJSON(msg)
        if err != nil {
            client.Close()
            delete(s.clients, client)
        }
    }

    return nil
}

func (s *ChatService) ProcessMessages(conn *websocket.Conn) {
    defer conn.Close()

    for {
        var msg domain.WebSocketExampleMessage
        err := conn.ReadJSON(&msg)
        if err != nil {
            s.UnregisterClient(conn)
            break
        }

        s.BroadcastMessage(msg)
    }
}
