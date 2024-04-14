package v1

import (
	"fmt"
	"net/http"

	"github.com/chatbot/internal/service"
	"github.com/chatbot/pkg/config"
	"github.com/chatbot/pkg/logger"

	"github.com/gin-gonic/gin"
)

// NewRouter ...
func NewRouter(handler *gin.Engine, l logger.Interface, s service.Services, cfg *config.Config) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.LoadHTMLGlob("templates/*")

	handler.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	handler.GET("/ping", func(c *gin.Context) {
		_, err := fmt.Fprint(c.Writer, "pong")
		if err != nil {
			return
		}
	})

	// Routers
	h := handler.Group("/api/v1")
	{
		newChatRoutes(h, s, l)
	}
}
