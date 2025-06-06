package main

import (
	"fmt"
	"sync"
)

// LogMsg represents a log message with a type and content
type LogMsg struct {
	Type    string
	Content string
}

func FanOut(source <-chan LogMsg) (<-chan LogMsg, <-chan LogMsg) {
	errors := make(chan LogMsg)
	infos := make(chan LogMsg)
	go func() {
		defer close(errors)
		defer close(infos)
		for msg := range source {
			switch msg.Type {
			case "error":
				errors <- msg
			case "info":
				infos <- msg
			}
		}
	}()
	return errors, infos
}

func FanOutExample() {
	source := make(chan LogMsg)
	errors, infos := FanOut(source)

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine to handle error messages
	go func() {
		defer wg.Done()
		for msg := range errors {
			fmt.Printf("Error: %s\n", msg.Content)
		}
	}()

	// Goroutine to handle info messages
	go func() {
		defer wg.Done()
		for msg := range infos {
			fmt.Printf("Info: %s\n", msg.Content)
		}
	}()

	// Send log messages to the source channel
	go func() {
		LogMsgs := []LogMsg{
			{"error", "Failed to connect to Redis"},
			{"info", "Process started successfully"},
			{"info", "New login from user"},
			{"error", "Database connection lost"},
			{"info", "User logged out"},
		}
		for _, msg := range LogMsgs {
			source <- msg
		}
		close(source)
	}()

	wg.Wait()
}
