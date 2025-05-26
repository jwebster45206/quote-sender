package notification

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// phoneRegex validates E.164 format: + followed by 1-15 digits
var phoneRegex = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

type SNSClient interface {
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

type SNSNotifier struct {
	client SNSClient
}

func NewSNSNotifier(client SNSClient) *SNSNotifier {
	return &SNSNotifier{
		client: client,
	}
}

func (s *SNSNotifier) Send(ctx context.Context, phoneNumber string, message string) error {
	if !phoneRegex.MatchString(phoneNumber) {
		return fmt.Errorf("invalid phone number format, must be in E.164 format (e.g., +1234567890)")
	}

	input := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	}

	_, err := s.client.Publish(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to publish message to SNS: %w", err)
	}

	return nil
}
