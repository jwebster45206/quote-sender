package ai

import (
	"context"
	"errors"
	"testing"
)

func TestMockProvider_GenerateQuote(t *testing.T) {
	tests := []struct {
		name      string
		mock      *MockProvider
		wantQuote string
		wantErr   error
		wantCalls int
	}{
		{
			name:      "returns quote",
			mock:      &MockProvider{},
			wantQuote: "Sharing tea with a fascinating stranger is one of life's true delights.",
			wantCalls: 1,
		},
		{
			name: "returns error",
			mock: &MockProvider{
				NextError: errors.New("foo"),
			},
			wantErr:   errors.New("foo"),
			wantCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.mock.GenerateQuote(context.Background())

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("GenerateQuote() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("GenerateQuote() unexpected error: %v", err)
				return
			}

			if got != tt.wantQuote {
				t.Errorf("GenerateQuote() = %v, want %v", got, tt.wantQuote)
			}

			if tt.mock.CallCount != tt.wantCalls {
				t.Errorf("GenerateQuote() called %d times, want %d", tt.mock.CallCount, tt.wantCalls)
			}
		})
	}
}
