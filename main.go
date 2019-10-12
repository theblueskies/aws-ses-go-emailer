package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/theblueskies/aws-ses-go-emailer/handler"
)

// PORT is the port where the service listens
const PORT = ":5000"

func main() {
	viper.BindEnv("SES_REGION")
	viper.BindEnv("SES_ACCESS_KEY")
	viper.BindEnv("SES_SECRET_KEY")
	viper.BindEnv("SENDER_EMAIL")

	sesRegion := viper.GetString("SES_REGION")
	sesAccessKey := viper.GetString("SES_ACCESS_KEY")
	sesSecretKey := viper.GetString("SES_SECRET_KEY")
	senderEmail := viper.GetString("SENDER_EMAIL")
	recipientEmail := viper.GetString("RECIPIENT_EMAIL")

	sesWorker := NewSESWorker(sesRegion, sesAccessKey, sesSecretKey, senderEmail, recipientEmail)

	// Set the "ENV" key in production to run in ReleaseMode
	if strings.ToLower(os.Getenv("ENV")) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := handler.GetRouter(sesWorker)
	r.Run(PORT) // listen and serve on 0.0.0.0:8080
}

// NewSESWorker returns a new instance of an SESWorker
func NewSESWorker(sesRegion, sesAccessKey, sesSecretKey, senderEmail, recipientEmail string) *handler.SESWorker {
	if sesAccessKey == "" || sesSecretKey == "" {
		err := fmt.Errorf("Access keys not set")
		panic(err)
	}

	if senderEmail == "" {
		panic(
			fmt.Errorf(
				"SENDER_EMAIL must be set. This is the verified sender email that is registered to send emails"))
	}

	if recipientEmail == "" {
		recipientEmail = senderEmail
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
		SenderEmail:    senderEmail,
		RecipientEmail: recipientEmail,
		Region:         sesRegion,
		AccessKey:      sesAccessKey,
		SecretKey:      sesSecretKey,
		Ses:            svc,
	}
}
