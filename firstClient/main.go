package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func sendFibonacciRequest(prev, next int) (int, error) {
	requestBody := map[string]int{"prev": prev, "next": next}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post("http://localhost:8080/calcFib", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Unknown state: %d", resp.StatusCode)
	}

	var result struct {
		Result int `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Result, nil
}

func consumeFromRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"fibonacci_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			var result map[string]int
			if err := json.Unmarshal(d.Body, &result); err != nil {
				log.Printf("Decoding error: %v", err)
				continue
			}
			fmt.Printf("The resulting calculation of the Fibonacci number: %d\n", result["result"])
		}
	}()
	select {}
}

func main() {

	go consumeFromRabbitMQ()

	prev, next := 0, 1
	for i := 0; i < 8; i++ {
		result, err := sendFibonacciRequest(prev, next)
		if err != nil {
			log.Printf("Error calculating Fibonacci: %v", err)
			continue
		}

		prev, next = next, result
	}
}
