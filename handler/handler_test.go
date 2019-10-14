package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/stretchr/testify/assert"
)

func TestHandleEmailSuccessFormData(t *testing.T) {
	mockSES := &mockSES{}
	s := &SESWorker{
		Region:    "us-east-1",
		AccessKey: "asdfa",
		SecretKey: "asdfa",
		Ses:       mockSES,
	}
	r := GetRouter(s)

	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("from", "email@sender.org")
	v.Add("subject", "the best subject")
	v.Add("body", "got a body")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/email", strings.NewReader(v.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(v.Encode())))
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
