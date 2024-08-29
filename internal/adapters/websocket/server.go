package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return nil, err
	}

	return conn, nil
}
