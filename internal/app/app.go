package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	v1 "github.com/chatbot/internal/api/route/v1"
	"github.com/chatbot/internal/service"
	"github.com/chatbot/pkg/config"
	"github.com/chatbot/pkg/logger"
	"github.com/chatbot/pkg/openai"
	"github.com/chatbot/pkg/server"
	"github.com/chatbot/pkg/storage"
)

// Start service
func Run(cfg *config.Config) {
	l := logger.New(cfg.App.Name+"> ", cfg.Log.Level)

	//s3 ...
	_, err := storage.NewS3Client(cfg.AWS, l)
	if err != nil {
		l.Fatal("[RUn] Failed to create s3 client, error: ", err)
	}

	s3 := storage.NewInMemoryStorage()

	// gptClient ...
	gptClient := openai.NewGPTClient(cfg.GPT)
	if err != nil {
		l.Fatal("[Run] Failed to create gpt client, error: ", err)
	}
	// Use cases
	//chatService ...
	chatService := service.NewChatService(gptClient, s3, l, cfg)

	// HTTP Server
	handler := gin.Default()

	services := service.Services{
		Chat: chatService,
	}
	v1.NewRouter(handler, l, services, cfg)
	httpServer := server.New(handler, l, server.Port(cfg.HTTP.Port))
	l.Info("[Run] Server started on port: ", cfg.HTTP.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("[Run]> signal " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("[Run]> httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("[Run]> httpServer.Shutdown: %w", err))
	}
}
