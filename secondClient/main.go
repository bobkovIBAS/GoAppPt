package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

type FibonacciRequest struct {
	Prev int `json:"prev"`
	Next int `json:"next"`
}

type FibonacciResponse struct {
	Result int `json:"result"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func calculateFibonacci(w http.ResponseWriter, r *http.Request) {
	var req FibonacciRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Incorrect request", http.StatusBadRequest)
		return
	}
	result := req.Prev + req.Next
	resp := FibonacciResponse{Result: result}

	publishToRabbitMQ(resp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func publishToRabbitMQ(resp FibonacciResponse) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect Rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"fibonacci_queue", // Имя очереди
		false,             // Устойчивая очередь
		false,             // Автоматическое удаление
		false,             // Эксклюзивная очередь
		false,             // Нет ожидания
		nil,               // Дополнительные аргументы
	)
	failOnError(err, "Failed to send to queue")
	body, err := json.Marshal(resp)
	failOnError(err, "Failed to encode response to JSON")
	err = ch.Publish(
		"",
		q.Name,
		false, // Обязательная отправка
		false, // Мандатное сообщение
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	failOnError(err, "Failed to publish message")

	log.Printf("Message has been sent: %s", body)
}

func main() {
	http.HandleFunc("/calcFib", calculateFibonacci)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//curl -X POST -H "Content-Type: application/json" -d "{\"prev\":1,\"next\":2}" http://localhost:8080/calcFib
