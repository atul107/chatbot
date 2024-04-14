package service

import (
	"context"
	"mime/multipart"

	"github.com/chatbot/internal/model"
	"github.com/chatbot/pkg/config"
	"github.com/chatbot/pkg/logger"
	"github.com/chatbot/pkg/openai"
	"github.com/chatbot/pkg/storage"
)

// ChatService defines the interface for the chatbot functionality
type ChatService struct {
	OpenAIClient  openai.OpenAIClient
	StorageClient storage.StorageClient
	logger        logger.Interface
	cfg           *config.Config
}

// NewChatService creates a new instance of the ChatService with necessary dependencies
func NewChatService(Client openai.OpenAIClient, storageClient storage.StorageClient, logger logger.Interface, cfg *config.Config) *ChatService {
	return &ChatService{
		OpenAIClient:  Client,
		StorageClient: storageClient,
		logger:        logger,
		cfg:           cfg,
	}
}

// HandleMessage handles chatbot messages
func (s *ChatService) HandleMessage(ctx context.Context, messageData model.MessageRequest) (string, error) {
	return s.OpenAIClient.SendMessage(messageData.Message)
}

// HandleUpload handles image uploads
func (s *ChatService) HandleUpload(ctx context.Context, file *multipart.FileHeader, userId string) (string, error) {
	return s.StorageClient.Upload(file, userId)
}

// RetrieveImage handles requests to retrieve images
func (s *ChatService) RetrieveImage(ctx context.Context, imageName string, userId string) ([]byte, error) {
	objectName := userId + "-" + imageName
	image, err := s.StorageClient.Retrieve(objectName)
	if err != nil {
		return nil, err
	}
	return image, err
}
