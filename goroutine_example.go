package main

import (
	"fmt"
	"time"
)

func GoroutineExample() {
	fmt.Println("Hello, World!")
	defer fmt.Println("Goodbye, World!")
	for i := 1; i <= 5; i++ {
		go fmt.Printf("Message %d\n", i)
	}
	time.Sleep(3 * time.Second)
}
