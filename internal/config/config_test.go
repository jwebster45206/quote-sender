package config

import (
	"context"
	"os"
	"testing"
)

func TestLoadApp(t *testing.T) {
	tests := []struct {
		name          string
		phoneNumbers  string
		expectedError bool
		expectedCount int
		cleanupEnvVar bool
	}{
		{
			name:          "single phone number",
			phoneNumbers:  "+1234567890",
			expectedCount: 1,
		},
		{
			name:          "empty entries are skipped",
			phoneNumbers:  "+1234567890,,+1987654321",
			expectedCount: 2,
		},
		{
			name:          "missing env var",
			cleanupEnvVar: true,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cleanupEnvVar {
				os.Unsetenv("RECIPIENT_PHONE_NUMBERS")
			} else {
				os.Setenv("RECIPIENT_PHONE_NUMBERS", tt.phoneNumbers)
			}

			cfg, err := LoadApp(context.Background())

			if tt.expectedError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got := len(cfg.RecipientPhoneNumbers); got != tt.expectedCount {
				t.Errorf("expected %d phone numbers, got %d", tt.expectedCount, got)
			}
		})
	}
}
