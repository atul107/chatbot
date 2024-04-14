package storage

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"sync"
)

// StorageClient is an interface for storage operations.
type StorageClient interface {
	Upload(*multipart.FileHeader, string) (string, error)
	Retrieve(string) ([]byte, error)
}

// InMemoryStorage implements the Client interface, storing files in memory.
type InMemoryStorage struct {
	files map[string][]byte
	mu    sync.RWMutex
}

// NewInMemoryStorage initializes and returns a new InMemoryStorage.
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		files: make(map[string][]byte),
	}
}

// Upload saves the given file data in memory and returns a unique identifier.
func (s *InMemoryStorage) Upload(file *multipart.FileHeader, userId string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Open the file
	uploadedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	objectName := userId + "-" + file.Filename
	if _, exists := s.files[objectName]; exists {
		return "", errors.New("file already exists")
	}

	// Read the file into memory
	data, err := io.ReadAll(uploadedFile)
	if err != nil {
		return "", err
	}

	// Store the file data
	s.files[objectName] = data

	return objectName, nil
}

// Retrieve fetches the file data for the given identifier.
func (s *InMemoryStorage) Retrieve(name string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, exists := s.files[name]
	if !exists {
		return nil, fmt.Errorf("file with id %s not found", name)
	}
	return data, nil
}
