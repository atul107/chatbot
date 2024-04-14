package main

import (
	"log"

	"github.com/chatbot/internal/app"
	"github.com/chatbot/pkg/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
