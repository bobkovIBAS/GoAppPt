package server

import (
	"encoding/json"
	"net/http"
)

type MessageBus interface {
	PublishMessage(queue string, message any) error
}

type QueryData struct {
	messageBus MessageBus
	calculator Calculator
}

func NewQuery(bus MessageBus, calc Calculator) *QueryData {
	return &QueryData{messageBus: bus, calculator: calc}
}

func (h *QueryData) CalculateFibonacci(w http.ResponseWriter, r *http.Request) {
	var req struct {
		N int `json:"n"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result := h.calculator.FibonacciCalculation(req.N)

	if err := h.messageBus.PublishMessage("fibonacci_queue", map[string]int{"result": result}); err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"result": result})
}
