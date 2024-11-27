package main

import (
	"fmt"
	"log"
	"os"
	"project/internal/client"
)

func main() {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080"
	}

	cl := client.NewFibonacciClient(apiURL)
	n := 25
	res, err := cl.SendFibonacciRequest(n)
	if err != nil {
		log.Fatal("Failed to calculate", err)
	}

	fmt.Printf("For the number %d, the Fibonacci result = %d", n, res)
}
