package notification

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockSNSClient is a mock implementation of the SNS client interface
type mockSNSClient struct {
	mock.Mock
}

func (m *mockSNSClient) Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	args := m.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sns.PublishOutput), args.Error(1)
}

func TestSNSNotifier(t *testing.T) {
	ctx := context.Background()

	t.Run("successful send", func(t *testing.T) {
		mockClient := new(mockSNSClient)
		notifier := NewSNSNotifier(mockClient)

		phoneNumber := "+12025550123"
		message := "Test message"

		// Expect Publish to be called with the correct parameters
		mockClient.On("Publish",
			mock.Anything,
			mock.MatchedBy(func(input *sns.PublishInput) bool {
				return *input.PhoneNumber == phoneNumber && *input.Message == message
			}),
			mock.Anything,
		).Return(&sns.PublishOutput{}, nil)

		err := notifier.Send(ctx, phoneNumber, message)

		assert.NoError(t, err)
		mockClient.AssertExpectations(t)
	})

	t.Run("invalid phone number format", func(t *testing.T) {
		mockClient := new(mockSNSClient)
		notifier := NewSNSNotifier(mockClient)

		invalidPhoneNumbers := []string{
			"1234567890",          // Missing +
			"+invalid",            // Non-numeric
			"+0234567890",         // Starting with 0
			"+1",                  // Too short
			"+123456789012345678", // Too long
		}

		for _, phone := range invalidPhoneNumbers {
			err := notifier.Send(context.Background(), phone, "Test message")
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "invalid phone number format")
		}

		// Ensure the mock was never called
		mockClient.AssertNotCalled(t, "Publish")
	})

	t.Run("sns error", func(t *testing.T) {
		mockClient := new(mockSNSClient)
		notifier := NewSNSNotifier(mockClient)

		phoneNumber := "+1234567890"
		message := "Test message"

		// Simulate an SNS error
		mockClient.On("Publish",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, assert.AnError)

		err := notifier.Send(context.Background(), phoneNumber, message)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to publish message to SNS")
		mockClient.AssertExpectations(t)
	})
}
