package notification

import "context"

// Notifier defines the interface for sending notifications
type Notifier interface {
	Send(ctx context.Context, phoneNumber string, message string) error
}
