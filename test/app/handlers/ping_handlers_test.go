package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/routes"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routes.GenerateRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}\n", w.Body.String())
}
