package notification

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/viniciusmtsantos/golang-ecommerce/configs"
)

type NotificationClient interface {
	SendSMS(phone string, message string) error
}

type notificationClient struct {
	config configs.ApplicationConfig
}

func NewNotificationClient(config configs.ApplicationConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}

// AWS SNS
func (c notificationClient) SendSMS(phone string, message string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(c.config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(c.config.AWSAccessKeyID, c.config.AWSAccessKeySecret, ""),
	})
	if err != nil {
		return errors.New("failed to create an aws session")
	}

	svc := sns.New(sess)

	params := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phone),
	}

	_, err = svc.Publish(params)
	if err != nil {
		return err
	}

	return nil
}
