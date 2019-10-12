package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	e := &Email{
		Name:    "some name",
		From:    "email@sender.org",
		Subject: "the best subject",
		Body:    "got a body",
	}

	s := &SESWorker{
		Region:    "us-east-1",
		AccessKey: "asdfa",
		SecretKey: "asdfa",
		Ses:       &mockSES{},
	}
	err := s.SendEmail(e)
	assert.Nil(t, err)
}

//@TODO: Write tests for the error cases
