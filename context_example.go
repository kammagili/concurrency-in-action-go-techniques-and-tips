package main

import (
	"context"
	"fmt"
	"time"
)

func ContextTimeoutExample() {
	ch := make(chan string)
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Operation timed out")
		}
	}
}

func ContextCancelExample() {
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println("Received:", msg)
			case <-ctx.Done():
				fmt.Println("Shutting down...")
				return
			}
		}
	}()
	ch <- "Hello"
	cancel()
	// Give some time for the goroutine to shut down
	time.Sleep(1 * time.Second)
	fmt.Println("End of the program")
}
