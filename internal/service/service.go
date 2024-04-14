package service

import (
	"context"
	"mime/multipart"

	"github.com/chatbot/internal/model"
)

type Services struct {
	Chat Chat
}

type (
	Chat interface {
		HandleMessage(context.Context, model.MessageRequest) (string, error)
		HandleUpload(context.Context, *multipart.FileHeader, string) (string, error)
		RetrieveImage(context.Context, string, string) ([]byte, error)
	}
)
