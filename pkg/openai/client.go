package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chatbot/pkg/config"
)

// OpenAIClient...
type OpenAIClient interface {
	SendMessage(string) (string, error)
}

// GPTClient represents a client for communicating with the OpenAI API.
type GPTClient struct {
	config config.GPT
	Client *http.Client
}

// NewGPTClient creates a new instance of Client with the provided config.
func NewGPTClient(gptConfig config.GPT) *GPTClient {
	return &GPTClient{
		config: gptConfig,
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

// SendMessage sends a message to the OpenAI API and returns the response.
func (c *GPTClient) SendMessage(prompt string) (string, error) {
	// Define the request payload
	payload := map[string]interface{}{
		"model": c.config.Model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": prompt},
		},
		"max_tokens": 50,
	}

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling payload: %w", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set necessary headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.OpenAIKey))

	// Send the request
	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to OpenAI API: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	var response struct {
		Choices []struct {
			Message struct{
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI API")
}
