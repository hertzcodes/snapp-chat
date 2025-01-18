package main

import (
	"flag"
	"log"
	"os"

	"github.com/hertzcodes/snapp-chat/server/config"
	"github.com/hertzcodes/snapp-chat/server/internal/api/handlers/http"
	"github.com/hertzcodes/snapp-chat/server/internal/app"
)

var configPath = flag.String("config", "config.json", "server configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.NewMustApp(c)
	log.Fatal(http.Run(appContainer, c.Server))
	
}
