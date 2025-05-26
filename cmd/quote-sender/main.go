package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/jwebster45206/quote-sender/internal/ai"
	"github.com/jwebster45206/quote-sender/internal/config"
	"github.com/jwebster45206/quote-sender/internal/notification"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	if err := run(context.Background()); err != nil {
		slog.Error("application error", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.LoadApp(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	slog.Info("configuration loaded",
		"recipients", len(cfg.RecipientPhoneNumbers),
		"ai_provider", cfg.AIProvider,
		"notification_provider", cfg.NotificationProvider)

	// Initialize AI provider
	var aiProvider ai.Provider
	switch cfg.AIProvider {
	case "openai":
		aiProvider = ai.NewOpenAIProvider(ai.OpenAIConfig{
			APIKey: cfg.OpenAIAPIKey,
			Model:  cfg.OpenAIModel,
		})
		slog.Info("using OpenAI provider", "model", cfg.OpenAIModel)
	default:
		aiProvider = ai.NewMockProvider()
		slog.Info("using mock AI provider")
	}

	// Initialize notification provider
	var notifier notification.Notifier
	switch cfg.NotificationProvider {
	case "sns":
		awsCfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(config.AWSRegion))
		if err != nil {
			return fmt.Errorf("failed to load AWS config: %w", err)
		}
		snsClient := sns.NewFromConfig(awsCfg)
		notifier = notification.NewSNSNotifier(snsClient)
		slog.Info("using SNS notification provider")
	default:
		notifier = notification.NewMockNotifier()
		slog.Info("using mock notification provider")
	}

	quote, err := aiProvider.GenerateQuote(ctx)
	if err != nil {
		return fmt.Errorf("failed to generate quote: %w", err)
	}
	slog.Info("quote generated", "quote", quote)

	for _, phone := range cfg.RecipientPhoneNumbers {
		if err := notifier.Send(ctx, phone, quote); err != nil {
			return fmt.Errorf("failed to send quote to %s: %w", phone, err)
		}
		slog.Info("quote sent", "phone", phone)
	}

	slog.Debug("application run completed successfully")
	return nil
}
