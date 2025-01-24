package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Connect(username string, room string) *websocket.Conn {
	url := "ws://localhost:8080/connect"

	headers := make(http.Header)

	headers.Add("username", username)
	headers.Add("room", room)

	conn, _, err := websocket.DefaultDialer.Dial(url, headers)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	fmt.Println("Joined room " + room)
	return conn
}
