package ai

import "context"

// Provider defines the interface for AI services that can generate quotes
type Provider interface {
	// GenerateQuote generates an inspirational quote in the style of Uncle Iroh
	// Returns the generated quote and any error encountered
	GenerateQuote(ctx context.Context) (string, error)
}
