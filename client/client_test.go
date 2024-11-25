package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var apiURL string

func TestSendFibonacciRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req map[string]int
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to request: %v", err)
		}
		if req["prev"] != 1 || req["next"] != 2 {
			t.Errorf(" %+v", req)
		}

		resp := map[string]int{"result": 3}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	originalURL := apiURL
	apiURL = server.URL
	defer func() { apiURL = originalURL }()

	result, err := sendFibonacciRequest(1, 2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != 3 {
		t.Errorf("Expected result 3, got %d", result)
	}
}
