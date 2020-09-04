package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	HealthCheck(w, r)
	resp := w.Result()
	assert.Equal(t, 204, resp.StatusCode, "expected 204 status code")
}
