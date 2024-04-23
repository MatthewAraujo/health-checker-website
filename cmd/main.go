package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/MatthewAraujo/health-checker-website/cmd/api"
	"github.com/MatthewAraujo/health-checker-website/cmd/config"
)

func main() {
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port))
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	slog.Info("Server started on port " + config.Envs.Port)
}
