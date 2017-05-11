package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubmitHandler(t *testing.T) {
	t.Run("TestIncorrectMethodError", func(t *testing.T) {
		handler := http.HandlerFunc(submitHandler)
		req, err := http.NewRequest(http.MethodGet, "/submit", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status code: %d, got: %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})
}
