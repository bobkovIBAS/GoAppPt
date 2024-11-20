package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type FibonacciRequest struct {
	Prev int `json:"prev"`
	Next int `json:"next"`
}

type FibonacciResponse struct {
	Result int `json:"result"`
}

func calculateFibonacci(w http.ResponseWriter, r *http.Request) {
	var req FibonacciRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Incorrect request", http.StatusBadRequest)
		return
	}

	result := req.Prev + req.Next
	resp := FibonacciResponse{Result: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/calcFib", calculateFibonacci)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
