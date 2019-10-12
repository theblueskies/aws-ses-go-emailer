package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/stretchr/testify/assert"
)

func TestHandleEmailSuccess(t *testing.T) {
	mockSES := &mockSES{}
	s := &SESWorker{
		Region:    "us-east-1",
		AccessKey: "asdfa",
		SecretKey: "asdfa",
		Ses:       mockSES,
	}
	r := GetRouter(s)
	e := Email{
		Name:    "some name",
		From:    "email@sender.org",
		Subject: "the best subject",
		Body:    "got a body",
	}

	eJSON, _ := json.Marshal(e)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/email", bytes.NewBuffer(eJSON))
	r.ServeHTTP(w, req)

	var b Response
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "success", b.Status)
	assert.Equal(t, true, mockSES.called)
	assert.Equal(t, 1, mockSES.callCount)
}

func TestHandleEmailFail(t *testing.T) {
	mockSES := &mockSES{}
	s := &SESWorker{
		Region:    "us-east-1",
		AccessKey: "asdfa",
		SecretKey: "asdfa",
		Ses:       mockSES,
	}
	r := GetRouter(s)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/email", nil)
	r.ServeHTTP(w, req)

	var b Response
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "error", b.Status)
	assert.Equal(t, false, mockSES.called)
}

func TestHealth(t *testing.T) {
	mockSES := &mockSES{}
	s := &SESWorker{
		Region:    "us-east-1",
		AccessKey: "asdfa",
		SecretKey: "asdfa",
		Ses:       mockSES,
	}
	r := GetRouter(s)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	var b Response
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ok", b.Status)
}

// mockSES is used to mock the SES object
type mockSES struct {
	sesiface.SESAPI
	output    *ses.SendEmailOutput
	err       error
	called    bool
	callCount int
}

// SendEmail is the implemented method from the interface sesiface.SESAPI for mockSES struct
func (m *mockSES) SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	m.called = true
	m.callCount++
	return m.output, m.err
}
