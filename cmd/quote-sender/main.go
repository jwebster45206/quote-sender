package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jwebster45206/quote-sender/internal/config"
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
	slog.Info("configuration loaded", "approvedPhones", len(cfg.ApprovedPhoneNumbers))

	// TODO: Initialize AI client
	// slog.Debug("initializing AI client")

	// TODO: Initialize SNS client
	// slog.Debug("initializing SNS client")

	// TODO: Generate quote
	// slog.Info("generating quote")

	// TODO: Send quote via SNS
	// slog.Info("sending quote")

	slog.Debug("application run completed successfully")
	return nil
}
