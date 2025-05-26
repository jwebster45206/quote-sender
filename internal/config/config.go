package config

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
)

type AppConfig struct {
	// phone numbers that can receive messages
	ApprovedPhoneNumbers []string
}

const (
	AWSRegion = "us-east-2"
)

// LoadApp creates a new AppConfig from environment variables
func LoadApp(ctx context.Context) (*AppConfig, error) {
	// Load AWS config to validate credentials
	// This will use the default credential chain:
	// - Environment variables (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)
	// - Shared credentials file (~/.aws/credentials)
	// - IAM role (if running on EC2 or Lambda)
	if _, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(AWSRegion),
	); err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	cfg := &AppConfig{}
	if phones := os.Getenv("APPROVED_PHONE_NUMBERS"); phones != "" {
		for _, phone := range strings.Split(phones, ",") {
			if trimmed := strings.TrimSpace(phone); trimmed != "" {
				cfg.ApprovedPhoneNumbers = append(cfg.ApprovedPhoneNumbers, trimmed)
			}
		}
	}
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	return cfg, nil
}

func (c *AppConfig) validate() error {
	if len(c.ApprovedPhoneNumbers) == 0 {
		return fmt.Errorf("APPROVED_PHONE_NUMBERS is required")
	}
	return nil
}
