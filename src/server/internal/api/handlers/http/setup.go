package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hertzcodes/snapp-chat/server/config"
	"github.com/hertzcodes/snapp-chat/server/internal/app"
	"github.com/nats-io/nats.go"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {

	// Connect to NATS server
	url := fmt.Sprintf("%s:%d", appContainer.Config().Nats.Host, appContainer.Config().Nats.Port)
	nc, err := nats.Connect(url)
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Println(err)
	}
	// Create a stream for chat messages
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "SnappChat",
		Subjects: []string{"SnappChat.*"},
	})
	if err != nil {
		log.Fatal("failed to create stream", err)
	}
	rooms := make(map[string]uint)
	http.HandleFunc("/login", Login(appContainer))
	http.HandleFunc("/connect", Connect(appContainer, js, rooms))
	// http.HandleFunc("/signup")
	connPath := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return http.ListenAndServe(connPath, nil)
}
