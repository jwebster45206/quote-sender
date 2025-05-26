package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	// phone numbers that will receive messages
	RecipientPhoneNumbers []string
	// AI provider to use (openai or mock)
	AIProvider string
	// OpenAI API key for the OpenAI provider
	OpenAIAPIKey string
	// OpenAI model to use (defaults to gpt-3.5-turbo if not set)
	OpenAIModel string
	// Notification provider to use (sns or mock)
	NotificationProvider string
}

const (
	AWSRegion = "us-east-2"
)

// LoadApp creates a new AppConfig from environment variables
func LoadApp(ctx context.Context) (*AppConfig, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// We don't return an error here as .env file is optional
		slog.Debug("no .env file found, using environment variables")
	}

	cfg := &AppConfig{}
	if phones := os.Getenv("RECIPIENT_PHONE_NUMBERS"); phones != "" {
		for _, phone := range strings.Split(phones, ",") {
			if trimmed := strings.TrimSpace(phone); trimmed != "" {
				cfg.RecipientPhoneNumbers = append(cfg.RecipientPhoneNumbers, trimmed)
			}
		}
	}

	cfg.AIProvider = strings.ToLower(os.Getenv("AI_PROVIDER"))
	if cfg.AIProvider == "" {
		cfg.AIProvider = "mock"
	}

	// Config OpenAI API if using OpenAI provider
	if cfg.AIProvider == "openai" {
		cfg.OpenAIAPIKey = os.Getenv("OPENAI_API_KEY")
		if cfg.OpenAIAPIKey == "" {
			return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required when using OpenAI provider")
		}
		cfg.OpenAIModel = os.Getenv("OPENAI_MODEL")
		if cfg.OpenAIModel == "" {
			cfg.OpenAIModel = "gpt-4o-mini"
		}
	}

	cfg.NotificationProvider = strings.ToLower(os.Getenv("NOTIFICATION_PROVIDER"))
	if cfg.NotificationProvider == "" {
		cfg.NotificationProvider = "mock"
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	return cfg, nil
}

func (c *AppConfig) validate() error {
	if len(c.RecipientPhoneNumbers) == 0 {
		return fmt.Errorf("RECIPIENT_PHONE_NUMBERS is required")
	}
	return nil
}
