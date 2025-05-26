package notification

import (
	"context"
	"fmt"
)

type MockNotifier struct {
	LastPhoneNumber string
	LastMessage     string
	ShouldError     bool
}

func NewMockNotifier() *MockNotifier {
	return &MockNotifier{
		ShouldError: false,
	}
}

func (m *MockNotifier) Send(ctx context.Context, phoneNumber string, message string) error {
	if m.ShouldError {
		return fmt.Errorf("mock notification error")
	}
	m.LastPhoneNumber = phoneNumber
	m.LastMessage = message
	return nil
}
