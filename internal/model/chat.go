package model

// MessageRequest represents a chat message
type MessageRequest struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

// MessageResponse ...
type MessageResponse struct {
	Message string `json:"message"`
}

// ImageUploadResponse ...
type ImageUploadResponse struct {
	Message string `json:"message"`
	ImageID string `json:"url"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

// ImageRetriveResponse ...
type ImageRetriveResponse struct {
	Image []byte `json:"image"`
}
