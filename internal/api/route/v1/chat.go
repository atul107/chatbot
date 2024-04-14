package v1

import (
	"github.com/chatbot/internal/api/controller"
	"github.com/chatbot/internal/api/middleware"
	"github.com/chatbot/internal/service"
	"github.com/chatbot/pkg/logger"

	"github.com/gin-gonic/gin"
)

// newChatRoutes ...
func newChatRoutes(handler *gin.RouterGroup, s service.Services, l logger.Interface) {
	r := &controller.ChatController{ChatService: s.Chat, Logger: l}

	// Define routes
	h := handler.Group("")
	{
		h.Use(middleware.Authenticate()).POST("/message", r.MessageHandler)
		h.Use(middleware.Authenticate()).POST("/upload", r.ImageUploadHandler)
		h.Use(middleware.Authenticate()).GET("/image/:user_id/:name", r.ImageRetriveHandler)
	}
}
