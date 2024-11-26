package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FibonacciClient struct {
	apiURL string
}

func NewFibonacciClient(apiURL string) *FibonacciClient {
	return &FibonacciClient{apiURL: apiURL}
}

func (c *FibonacciClient) SendFibonacciRequest(n int) (int, error) {
	req := map[string]int{"n": n}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(c.apiURL+"/calcFib", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Result int `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Result, nil
}
