package main

import (
	"fmt"
	"log"
	"project/internal/client"
)

func main() {
	cl := client.NewFibonacciClient("http://localhost:8080")

	n := 25
	res, err := cl.SendFibonacciRequest(n)
	if err != nil {
		log.Fatal("Failed to calculate", err)
	}

	fmt.Printf("For the number %d, the Fibonacci result = %d", n, res)
}
