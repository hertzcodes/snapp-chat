package http

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hertzcodes/snapp-chat/server/internal/app"
	tt "github.com/hertzcodes/snapp-chat/server/internal/time"
	"github.com/nats-io/nats.go"
)

type Message struct {
	Data []byte
	User string
	Time string
}

var (
	ErrMessageNotSent        = errors.New("message not sent")
	AnnMessageNotSent string = "\n[%s][SERVER] Error sending message\n"
	AnnJoinedChat     string = "\n[%s][SERVER] %s HAS JOINED THE CHAT\n"
	AnnLeftChat       string = "\n[%s][SERVER] %s HAS LEFT THE CHAT\n"
	AnnNewDay         string = "\n	%s\n"
)

func Connect(appContainer app.App, js nats.JetStreamContext, rooms map[string]uint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		conn, err := Upgrade(w, r) // upgrades to a websocket connection
		if err != nil {
			log.Println("Error upgrading connection:", err)
			return
		}
		defer conn.Close()
		username := r.Header.Get("username")
		room := r.Header.Get("room")

		rooms[room]++
		defer func() {
			fmt.Println(rooms)
			rooms[room]--
			fmt.Println(rooms)
		}()
		// subscribes to a jetstream subject
		sub, err := js.Subscribe(fmt.Sprintf("SnappChat.%s", room), func(m *nats.Msg) {

			var d Message

			decodeMessage(m.Data, &d)

			switch {
			case d.User == "server":
				if err := conn.WriteMessage(websocket.TextMessage, d.Data); err != nil {
					log.Println("Error sending message to WebSocket:", err)
				}
			default:
				msg := fmt.Sprintf("[%s] %s: %s", d.Time, d.User, string(d.Data))
				// sends the message to all clients in room
				if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
					log.Println("Error sending message to WebSocket:", err)
				}
			}

		}, nats.AckAll())
		if err != nil {
			log.Fatal("Error subscribing to NATS:", err)
			return
		}

		// reading connection messages and publish
		go func() {
			publish(js, room, Message{User: "server", Data: []byte(fmt.Sprintf(AnnJoinedChat, tt.GetTime(true), username))})

			defer sub.Unsubscribe() // unsubscribe if function dies

			go func() {
				for {
					if t := time.Now(); t.Hour() == 0 && t.Minute() == 0 && t.Second() <= 1 {
						publish(js, room, Message{User: "server", Data: []byte(fmt.Sprintf(AnnNewDay, tt.GetDay(true)))})
						time.Sleep(time.Second * 1)
					} // seperates each day
				}
			}()

			for {

				messageType, msg, err := conn.ReadMessage()

				if err != nil {
					publish(js, room, Message{User: "server", Data: []byte(fmt.Sprintf(AnnLeftChat, tt.GetTime(true), username))})
					return
				} // announces that user has left the chat

				switch string(msg) {

				case "#users":
					conn.WriteMessage(messageType, []byte(fmt.Sprintf("Online people count: %d", rooms[room])))
				default:
					if err := publish(js, room, Message{User: username, Time: tt.GetTime(true), Data: msg}); err != nil {
						conn.WriteMessage(messageType, []byte("failed to send message!")) // shows the error to sender
						log.Println(err)

					}
				}
			}
		}()
		select {} // I know...
	}
}

func publish(ctx nats.JetStreamContext, room string, d Message) error {

	msg, err := encodeMessage(d)
	if err != nil {
		return ErrMessageNotSent
	}

	_, err = ctx.Publish(fmt.Sprintf("SnappChat.%s", room), msg)
	return err

}

func encodeMessage(msg Message) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(msg)
	return buf.Bytes(), err
}

func decodeMessage(data []byte, msg *Message) {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	decoder.Decode(msg)
}
