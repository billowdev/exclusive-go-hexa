package services

import (
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/websocket"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/websocket"
	"github.com/gorilla/websocket"
)

type WebSocketMessageService struct {
	conn *websocket.Conn
}

func NewWebSocketMessageService(conn *websocket.Conn) ports.IWebSocketExamplePorts {
	return &WebSocketMessageService{conn: conn}
}

// ReceiveMessage implements ports.IWebSocketExamplePorts.
func (w *WebSocketMessageService) ReceiveMessage() (domain.WebSocketExampleMessage, error) {
	var msg domain.WebSocketExampleMessage
	err := w.conn.ReadJSON(&msg)
	return msg, err
}

// SendMessage implements ports.IWebSocketExamplePorts.
func (w *WebSocketMessageService) SendMessage(msg domain.WebSocketExampleMessage) error {
	return w.conn.WriteJSON(msg)
}

func (s *WebSocketMessageService) ProcessMessage() error {
	msg, err := s.ReceiveMessage()
	if err != nil {
		return err
	}

	// Business logic could be added here
	// For now, we'll just echo the message back.
	return s.SendMessage(msg)
}
