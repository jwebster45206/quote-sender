# Security Considerations

## Credentials

This application uses the AWS SDK's default credential chain. No credentials are stored in the codebase.

### Local Development
Credentials are loaded from one of:
- Environment variables (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)
- Shared credentials file (~/.aws/credentials)
- IAM instance role (if running on EC2)

### Lambda Deployment
When deployed to Lambda, the function will use the Lambda execution role's permissions.

## Phone Numbers

Approved phone numbers are loaded from environment variables. No phone numbers are stored in the codebase.

### Environment Variables
- `APPROVED_PHONE_NUMBERS`: Comma-separated list of phone numbers that can receive messages

## Best Practices

1. Never commit AWS credentials to the repository
2. Never commit real phone numbers to the repository
3. Use example phone numbers (e.g., "+1234567890") in documentation and tests
4. Keep the approved phone number list minimal and regularly review it
5. When deploying to Lambda, follow the principle of least privilege when configuring the IAM role
