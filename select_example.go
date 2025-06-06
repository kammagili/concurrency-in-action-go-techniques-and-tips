package main

import (
	"fmt"
	"time"
)

func TimeoutExample() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Data received"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout occurred")
	}
}

func GracefulShutdownExample() {
	ch := make(chan string)
	done := make(chan bool)
	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println("Received:", msg)
			case <-done:
				fmt.Println("Shutting down...")
				return
			}
		}
	}()
	ch <- "Hello"
	close(done)
	time.Sleep(1 * time.Second)
}
