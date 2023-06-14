package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_specialHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/foo?item=apple", nil)
	w := httptest.NewRecorder()

	specialHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code. Actual status code %v", resp.StatusCode)
	}
}
