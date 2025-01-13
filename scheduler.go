package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Set the number of logical processors
	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup

	// Create a large number of goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Simulate some work
			time.Sleep(time.Millisecond * time.Duration(id%10))

			threadID, err := getThreadID()
			if err != nil {
				fmt.Printf("Error getting thread ID for goroutine %d: %v\n", id, err)
				return
			}
			fmt.Printf("Goroutine %d on thread %d\n", id, threadID)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print scheduler statistics
	fmt.Printf("\nScheduler Statistics:\n")
	fmt.Printf("Logical Processors: %d\n", runtime.GOMAXPROCS(-1))
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
}

func getThreadID() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	var id int
	_, err := fmt.Sscanf(string(buf[:n]), "goroutine %d ", &id)
	if err != nil {
		return 0, fmt.Errorf("failed to parse goroutine ID: %w", err)
	}
	return id, nil
}
