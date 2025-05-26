# quote-sender

Generates an inspirational quote from Uncle Iroh, and sends it via SMS to an approved recipient.

## Overview

- V1: Runs locally on developer workstation, sends via SNS
- V2: Runs as AWS Lambda function (planned)

## Configuration

### Environment Variables

- `APPROVED_PHONE_NUMBERS`: Comma-separated list of phone numbers that can receive messages
- AWS credentials (any of):
  - Environment: `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`
  - Shared credentials file: `~/.aws/credentials`
  - IAM role (when running on EC2 or Lambda)

### AWS Configuration

- Region: us-east-2 (hardcoded)
- Required permissions: SNS:Publish to send messages
- When running in Lambda (V2), uses the Lambda execution role

## Security Considerations

- Uses AWS SDK's default credential chain
- Only sends to pre-approved phone numbers
- Phone number list managed via environment variables
- Validates all recipient numbers against approved list

## Design Notes

- Supports multiple AI providers via configuration
- Uses AWS SNS for reliable SMS delivery
- JSON structured logging for monitoring

### Future Considerations
- Consider tracking used quotes to avoid repetition
- Lambda deployment (V2)
- Expanded AI provider options

## Local Development

1. Configure AWS credentials (see Configuration section)
2. Set approved phone numbers:
   ```sh
   export APPROVED_PHONE_NUMBERS="+1234567890,+1234567891"
   ```
3. Run the application:
   ```sh
   go run ./cmd/quote-sender
   ```