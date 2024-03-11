package main

import "fmt"

func sendMessages(messages chan<- string, texts []string) {
	for _, text := range texts {
		messages <- text // Sending value to the channel
	}
	close(messages) // Closing the channel after sending all messages
}

func readMessages(messages <-chan string) {
	for message := range messages {
		fmt.Println("Received:", message) // Receiving value from the channel
	}
}

func main() {
	messages := make(chan string)

	// Send-only channel
	go sendMessages(messages, []string{"Hello", "World"})

	// Receive-only channel
	readMessages(messages)
}
