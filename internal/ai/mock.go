package ai

import "context"

// MockProvider implements Provider interface for testing purposes
type MockProvider struct {
	// NextError is the error that will be returned by the next call to GenerateQuote
	NextError error
	// CallCount tracks the number of times GenerateQuote was called
	CallCount int
}

func NewMockProvider() Provider {
	return &MockProvider{}
}

// GenerateQuote implements Provider interface
func (m *MockProvider) GenerateQuote(ctx context.Context) (string, error) {
	m.CallCount++
	if m.NextError != nil {
		return "", m.NextError
	}
	return "Sharing tea with a fascinating stranger is one of life's true delights.", nil
}
