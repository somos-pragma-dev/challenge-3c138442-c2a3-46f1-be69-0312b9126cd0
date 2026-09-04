package tests

import (
	"internal/app/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	r := gin.Default()
	accountHandler := handlers.NewAccountHandler()
	accountHandler.RegisterRoutes(r)
	req, _ := http.NewRequest("POST", "/accounts", strings.NewReader(`{"number":"123456789","balance":1000.0,"owner":"John Doe"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code!= http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, w.Code)
	}
}