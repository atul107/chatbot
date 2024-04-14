package tests

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"

	"github.com/chatbot/internal/api/controller"
	"github.com/chatbot/internal/api/middleware"
	"github.com/chatbot/internal/service"
	"github.com/chatbot/pkg/config"
	"github.com/chatbot/pkg/logger"
	"github.com/chatbot/pkg/openai"
	"github.com/chatbot/pkg/storage"
)

func initConfig() (*config.Config, error) {
	cfg := &config.Config{}

	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf(".env file not found or error loading it: %v. Relying on system environment variables", err)
	}
	return cfg, nil
}

func setupRouter() *gin.Engine {
	cfg, _ := initConfig()
	router := gin.Default()

	l := logger.New(cfg.App.Name+"> ", cfg.Log.Level)

	//s3 ...
	_, err := storage.NewS3Client(cfg.AWS, l)
	if err != nil {
		l.Fatal("[RUn] Failed to create s3 client, error: ", err)
	}

	s3 := storage.NewInMemoryStorage()

	gptClient := openai.NewGPTClient(cfg.GPT)
	if err != nil {
		l.Fatal("[Run] Failed to create gpt client, error: ", err)
	}
	chatService := service.NewChatService(gptClient, s3, l, cfg)

	r := &controller.ChatController{ChatService: chatService, Logger: l}

	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(middleware.Authenticate())
	protectedRoutes.POST("/message", r.MessageHandler)
	protectedRoutes.POST("/upload", r.ImageUploadHandler)
	protectedRoutes.GET("/image/:user_id/:name", r.ImageRetriveHandler)

	return router
}

func performRequest(router http.Handler, method, path string, body []byte, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+"fgcfgjfcgjcfghjf")
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}
