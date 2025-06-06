package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Stage 1: Read lines from a file
func readLines(filename string, out chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		close(out)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out <- scanner.Text()
	}
	close(out)
}

// Stage 2: Count words in each line
func countWords(in <-chan string, out chan<- int) {
	for line := range in {
		out <- len(strings.Fields(line))
	}
	close(out)
}

// Stage 3: Print word counts
func printCounts(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for count := range in {
		fmt.Println("Word count:", count)
	}
}

func PipelineExample() {
	lines := make(chan string)
	wordCounts := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go readLines("example.txt", lines)
	go countWords(lines, wordCounts)
	go printCounts(wordCounts, &wg)

	wg.Wait()
}
