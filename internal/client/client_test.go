package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendFibonacciRequest(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST; got %s", r.Method)
		}
		if r.URL.Path != "/calcFib" {
			t.Errorf("Expected URL path /calcFib; got %s", r.URL.Path)
		}

		var reqBody map[string]int
		json.NewDecoder(r.Body).Decode(&reqBody)
		defer r.Body.Close()

		n, ok := reqBody["n"]
		if !ok || n != 5 {
			t.Errorf("Expected 'n' to be 5; got %d", n)
		}

		respBody, _ := json.Marshal(map[string]int{"result": 8})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer testServer.Close()

	client := NewFibonacciClient(testServer.URL)
	result, err := client.SendFibonacciRequest(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("Expected result 8; got %d", result)
	}
}
