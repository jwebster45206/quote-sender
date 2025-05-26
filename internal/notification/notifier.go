package notification

// Notifier defines the interface for sending notifications
type Notifier interface {
	Send(phoneNumber string, message string) error
}
