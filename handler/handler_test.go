package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleEmailSuccess(t *testing.T) {
	r := GetRouter()
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

}

func TestHandleEmailFail(t *testing.T) {
	r := GetRouter()

	// eJSON, _ := json.Marshal(e)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/email", nil)
	r.ServeHTTP(w, req)

	var b Response
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "error", b.Status)

}

func TestHealth(t *testing.T) {
	r := GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	var b Response
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ok", b.Status)
}
