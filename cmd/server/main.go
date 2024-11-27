package main

import (
	"log"
	"net/http"
	"os"
	"project/internal/server"
)

func main() {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		rabbitmqURL = "amqp://guest:guest@localhost:5672/"
	}

	con, err := server.ConnectingToRabbitmq(rabbitmqURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer con.Conn.Close()

	handler := server.NewQuery(con, server.FibonacciCalculatorData{})

	http.HandleFunc("/calcFib", handler.CalculateFibonacci)

	log.Println("Server stated ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
