package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chatbot/internal/model"
)

func TestMessageHandler(t *testing.T) {
	// Setup
	router := setupRouter()
	reqBody := model.MessageRequest{UserID: 1, Message: "What is trending in fashion industry?"}
	requestBody, _ := json.Marshal(reqBody)
	w := performRequest(router, "POST", "/api/v1/message", requestBody, true)

	// Assertion
	assert.Equal(t, http.StatusOK, w.Code)

	var response model.MessageResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.Message)
}
