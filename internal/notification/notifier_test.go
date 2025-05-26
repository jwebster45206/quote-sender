package notification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockNotifier(t *testing.T) {
	t.Run("successful notification", func(t *testing.T) {
		notifier := NewMockNotifier()
		phoneNumber := "+1234567890"
		message := "Test message"

		err := notifier.Send(phoneNumber, message)

		assert.NoError(t, err)
		assert.Equal(t, phoneNumber, notifier.LastPhoneNumber)
		assert.Equal(t, message, notifier.LastMessage)
	})
	t.Run("error case", func(t *testing.T) {
		notifier := NewMockNotifier()
		notifier.ShouldError = true
		err := notifier.Send("+1234567890", "Test message")
		assert.Error(t, err)
		assert.Equal(t, "mock notification error", err.Error())
	})
}
