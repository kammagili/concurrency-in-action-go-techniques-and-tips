package main

import (
	"fmt"
	"sync"
	"time"
)

var maxGoroutines = 3
var semaphore = make(chan struct{}, maxGoroutines)
var wg = &sync.WaitGroup{}

func SemaphoreExample() {
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			printWithDelay(fmt.Sprintf("Processing task %d", i))
		}(i)
	}
	// Wait for all goroutines to finish
	wg.Wait()
}
func printWithDelay(message string) {
	defer wg.Done()
	// Acquire a semaphore slot
	semaphore <- struct{}{}
	defer func() { <-semaphore }()
	fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
	time.Sleep(3 * time.Second)
}
