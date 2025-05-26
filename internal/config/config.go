package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	// phone numbers that will receive messages
	RecipientPhoneNumbers []string

	// AI Provider config
	AIProvider   string
	OpenAIAPIKey string
	OpenAIModel  string

	// Notification Provider config
	NotificationProvider string
	SNSTopicARN          string
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

	// Config SNS if using SNS provider
	if cfg.NotificationProvider == "sns" {
		cfg.SNSTopicARN = os.Getenv("SNS_TOPIC_ARN")
		if cfg.SNSTopicARN == "" {
			return nil, fmt.Errorf("SNS_TOPIC_ARN environment variable is required when using SNS provider")
		}
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	return cfg, nil
}

func (c *AppConfig) validate() error {
	if len(c.RecipientPhoneNumbers) == 0 {
		return fmt.Errorf("APPROVED_PHONE_NUMBERS is required")
	}
	return nil
}
