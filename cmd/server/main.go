package main

import (
	"log"
	"net/http"
	"project/internal/server"
)

func main() {
	con, err := server.ConnectingToRabbitmq("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer con.Conn.Close()

	handler := server.NewQuery(con, server.FibonacciCalculatorData{})

	http.HandleFunc("/calcFib", handler.CalculateFibonacci)

	log.Println("Server stated ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
