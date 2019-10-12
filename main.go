package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/spf13/viper"
	"github.com/theblueskies/aws-ses-go-emailer/handler"
)

func main() {
	sesRegion := viper.GetString("SES_REGION")
	sesAccessKey := viper.GetString("SES_ACCESS_KEY")
	sesSecretKey := viper.GetString("SES_SECRET_KEY")
	senderEmail := viper.GetString("SENDER_EMAIL")

	sesWorker := NewSESWorker(sesRegion, sesAccessKey, sesSecretKey, senderEmail)

	r := handler.GetRouter(sesWorker)
	r.Run() // listen and serve on 0.0.0.0:8080
}

// NewSESWorker returns a new instance of an SESWorker
func NewSESWorker(sesRegion, sesAccessKey, sesSecretKey, senderEmail string) *handler.SESWorker {
	if sesAccessKey == "" || sesSecretKey == "" {
		err := fmt.Errorf("Access keys not set")
		panic(err)
	}

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(sesRegion),
		Credentials: credentials.NewStaticCredentials(sesAccessKey, sesSecretKey, ""),
	})

	if err != nil {
		err := fmt.Errorf("Error getting aws session")
		panic(err)
	}

	svc := ses.New(session)
	return &handler.SESWorker{
		SenderEmail: senderEmail,
		Region:      sesRegion,
		AccessKey:   sesAccessKey,
		SecretKey:   sesSecretKey,
		Ses:         svc,
	}
}
