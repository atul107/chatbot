package controller

import (
	"net/http"

	"github.com/chatbot/internal/model"
	"github.com/chatbot/internal/service"
	"github.com/chatbot/pkg/logger"
	"github.com/gin-gonic/gin"
)

// ChatController ...
type ChatController struct {
	ChatService service.Chat
	Logger      logger.Interface
}

// GetChatHandler ...
func (cc *ChatController) MessageHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var data model.MessageRequest
	if err := validateBody(c, &data); err != nil {
		cc.Logger.Error("[ChatController.MessageHandler] failed to parse body, error: ", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := cc.ChatService.HandleMessage(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Failed to process message. Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageResponse{Message: response})
}

// SendChatHandler ...
func (cc *ChatController) ImageUploadHandler(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.PostForm("user_id")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid file upload"})
		return
	}

	imageId, err := cc.ChatService.HandleUpload(ctx, file, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Failed to upload image"})
		return
	}
	c.JSON(http.StatusOK, model.ImageUploadResponse{Message: "Image uploaded successfully", ImageID: imageId})
}

// SendChatHandler ...
func (cc *ChatController) ImageRetriveHandler(c *gin.Context) {
	ctx := c.Request.Context()
	imageName := c.Param("name")
	userId := c.Param("user_id")

	image, err := cc.ChatService.RetrieveImage(ctx, imageName, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Failed to retrieve image"})
		return
	}
	c.JSON(http.StatusOK, model.ImageRetriveResponse{Image: image})
}
