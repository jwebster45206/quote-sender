package notification

import (
	"context"
	"fmt"
)

// MockNotifier is a mock implementation of the Notifier interface
type MockNotifier struct {
	LastPhoneNumber string
	LastMessage     string
	ShouldError     bool
}

// NewMockNotifier creates a new instance of MockNotifier
func NewMockNotifier() *MockNotifier {
	return &MockNotifier{
		ShouldError: false,
	}
}

// Send implements the Notifier interface for MockNotifier
func (m *MockNotifier) Send(ctx context.Context, phoneNumber string, message string) error {
	if m.ShouldError {
		return fmt.Errorf("mock notification error")
	}
	m.LastPhoneNumber = phoneNumber
	m.LastMessage = message
	return nil
}
