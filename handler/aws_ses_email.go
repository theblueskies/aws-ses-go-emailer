package handler

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

const (
	ACCESS_KEY = "AKIA22BVYOS7LVPNHZ6F"
	SECRET_KEY = "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"
)

// SESWorker is used to send the actual email. It implements the EmailStore
type SESWorker struct {
	SenderEmail string
	Region      string
	AccessKey   string
	SecretKey   string
	Ses         sesiface.SESAPI
}

// SendEmail is a wrapper around the AWS SES object and calls the SES.SendEmail method
func (s *SESWorker) SendEmail(e *Email) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(e.From),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				// @TODO: For pretty HTML emails, fill up this block with the HTML
				// block in "Data" below

				// Html: &ses.Content{
				// 	Charset: aws.String("UTF-8"),
				// 	Data:    aws.String("This message body contains HTML formatting. It can, for example, contain links like this one: <a class=\"ulink\" href=\"http://docs.aws.amazon.com/ses/latest/DeveloperGuide\" target=\"_blank\">Amazon SES Developer Guide</a>."),
				// },
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(e.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(e.Subject),
			},
		},
		Source: aws.String("bladerunneraws@gmail.com"),
	}

	result, err := s.Ses.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeConfigurationSetSendingPausedException:
				fmt.Println(ses.ErrCodeConfigurationSetSendingPausedException, aerr.Error())
			case ses.ErrCodeAccountSendingPausedException:
				fmt.Println(ses.ErrCodeAccountSendingPausedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return err
	}

	fmt.Println(result)
	return err
}
