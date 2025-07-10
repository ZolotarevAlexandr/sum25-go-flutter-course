package models

import (
	"errors"
	"time"
)

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidContent  = errors.New("invalid content")
)

// Message represents a chat message
type Message struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *Message) Validate() error {
	if err := ValidateUsername(m.Username); err != nil {
		return err
	}
	if err := ValidateContent(m.Content); err != nil {
		return err
	}
	return nil
}

// CreateMessageRequest represents the request to create a new message
type CreateMessageRequest struct {
	Username string `json:"username" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

// UpdateMessageRequest represents the request to update a message
type UpdateMessageRequest struct {
	Content string `json:"content" validate:"required"`
}

// HTTPStatusResponse represents the response for HTTP status code endpoint
type HTTPStatusResponse struct {
	StatusCode  int    `json:"status_code"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// NewMessage creates a new message with the current timestamp
func NewMessage(id int, username, content string) *Message {
	return &Message{
		ID:        id,
		Username:  username,
		Content:   content,
		Timestamp: time.Now(),
	}
}

func ValidateUsername(username string) error {
	if username == "" {
		return ErrInvalidUsername
	}
	return nil
}

func ValidateContent(content string) error {
	if content == "" {
		return ErrInvalidContent
	}
	return nil
}

// Validate checks if the create message request is valid
func (r *CreateMessageRequest) Validate() error {
	// Check if Username is not empty
	// Check if Content is not empty
	// Return appropriate error messages
	if err := ValidateUsername(r.Username); err != nil {
		return err
	}
	if err := ValidateContent(r.Content); err != nil {
		return err
	}
	return nil
}

// Validate checks if the update message request is valid
func (r *UpdateMessageRequest) Validate() error {
	// Check if Content is not empty
	// Return appropriate error messages
	if err := ValidateContent(r.Content); err != nil {
		return err
	}
	return nil
}
