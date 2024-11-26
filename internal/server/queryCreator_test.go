package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockMessageBus struct {
	publishedMessages []interface{}
	err               error
}

func (m *MockMessageBus) PublishMessage(queue string, message any) error {
	if m.err != nil {
		return m.err
	}
	m.publishedMessages = append(m.publishedMessages, message)
	return nil
}

type MockCalculator struct {
	expectedInput int
	result        int
}

func (m *MockCalculator) FibonacciCalculation(n int) int {
	if n != m.expectedInput {
		panic("Unexpected input to FibonacciCalculation")
	}
	return m.result
}

func TestCalculateFibonacciHandler(t *testing.T) {
	inputN := 5
	expectedResult := 8

	mockBus := &MockMessageBus{}
	mockCalculator := &MockCalculator{
		expectedInput: inputN,
		result:        expectedResult,
	}

	query := NewQuery(mockBus, mockCalculator)
	reqBody, _ := json.Marshal(map[string]int{"n": inputN})
	req := httptest.NewRequest(http.MethodPost, "/calcFib", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	query.CalculateFibonacci(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Status code = %d; expected %d", w.Result().StatusCode, http.StatusOK)
	}

	var respBody map[string]int
	err := json.NewDecoder(w.Body).Decode(&respBody)
	if err != nil {
		t.Fatalf("Error decoding the response: %v", err)
	}
	if respBody["result"] != expectedResult {
		t.Errorf("The result of the response = %d; expected %d", respBody["result"], expectedResult)
	}
	if len(mockBus.publishedMessages) != 1 {
		t.Errorf("1 published message is expected; received %d", len(mockBus.publishedMessages))
	} else {
		publishedMessage, ok := mockBus.publishedMessages[0].(map[string]int)
		if !ok {
			t.Errorf("The published message has the wrong type")
		} else if publishedMessage["result"] != expectedResult {
			t.Errorf("The result of the published message = %d; expected %d", publishedMessage["result"], expectedResult)
		}
	}
}
