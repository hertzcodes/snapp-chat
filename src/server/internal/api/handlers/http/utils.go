package http

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn , error) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return nil , err 
	}

	return conn , err 
}
