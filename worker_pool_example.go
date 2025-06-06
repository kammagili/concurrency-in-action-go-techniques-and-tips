package main

import (
	"fmt"
	"math/rand/v2"
)

// Worker function that processes tasks
func worker(id int, tasks <-chan int, results chan<- string) {
	for _ = range tasks {
		a := rand.IntN(1000) + rand.IntN(1000)
		results <- fmt.Sprintf("Result: %d, worker %d", a, id)
	}
}

func WorkerPoolExample(numWorkers int, numTasks int) {
	// Channels for tasks and results
	tasks := make(chan int, numTasks)
	results := make(chan string, numTasks)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results)
	}

	// Send tasks to the task channel
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)

	// Collect results
	var i int
	for i = 1; i <= numTasks; i++ {
		<-results
	}
}
