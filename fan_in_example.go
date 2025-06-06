package main

import (
	"fmt"
	"sync"
	"time"
)

// fanInWorker simulates a task producing results onto a channel for the fan-in example.
func fanInWorker(id int, jobs <-chan int, results chan<- string) {
	defer close(results) // Close the results channel when this worker is done
	for j := range jobs {
		fmt.Printf("FanInWorker %d: Started job %d\n", id, j)
		time.Sleep(time.Millisecond * 500) // Simulate work
		results <- fmt.Sprintf("FanInWorker %d finished job %d", id, j)
	}
}

func fanIn(sources ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	wg.Add(len(sources))
	for _, c := range sources {
		go func(c <-chan string) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func FanInExample() {
	numJobs := 5
	numWorkers := 3
	jobs := make(chan int, numJobs)

	// Create channels for each worker's results
	resultsChans := make([]<-chan string, numWorkers) // Slice of read-only channels for fanIn
	workerOutputs := make([]chan string, numWorkers)  // Slice of writeable channels for workers

	for w := 0; w < numWorkers; w++ {
		// Create a channel for this specific worker
		ch := make(chan string)
		workerOutputs[w] = ch
		resultsChans[w] = ch                        // Add the read-only version to the slice for fanIn
		go fanInWorker(w+1, jobs, workerOutputs[w]) // Use the renamed worker function
	}

	// Send jobs to the workers
	fmt.Println("Sending jobs...")
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to signal workers no more jobs are coming
	fmt.Println("All jobs sent.")

	// Use fanIn to merge results from all worker channels
	mergedResults := fanIn(resultsChans...)

	// Collect results from the merged channel
	fmt.Println("Waiting for results...")
	for res := range mergedResults {
		fmt.Printf("Result received: %s\n", res)
	}

	fmt.Println("All results received. Fan-in complete.")
}
