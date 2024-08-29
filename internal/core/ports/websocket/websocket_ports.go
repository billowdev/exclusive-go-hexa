package ports

import (
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/websocket"
	"github.com/gorilla/websocket"
)

type IWebSocketExamplePorts interface {
	SendMessage(msg domain.WebSocketExampleMessage) error
	ReceiveMessage() (domain.WebSocketExampleMessage, error)
	GetConnection() *websocket.Conn
}
