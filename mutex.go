package main

import (
	"Basics/utils"
	"fmt"
	"sync"
	"time"
)

// Counter is a struct that holds a count and a mutex
type Counter struct {
	count int
	mutex sync.Mutex
}

// Increment method safely increases the count
func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

// GetCount method safely retrieves the current count
func (c *Counter) GetCount() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func main() {
	utils.PrintSection("1. Without Mutex (Race Condition)")
	counterWithoutMutex := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterWithoutMutex++
		}()
	}
	wg.Wait()
	fmt.Printf("Counter without mutex: %d\n", counterWithoutMutex)

	utils.PrintSection("2. With Mutex (Race Condition Prevented)")
	counterWithMutex := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterWithMutex.Increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter with mutex: %d\n", counterWithMutex.GetCount())

	utils.PrintSection("3. Demonstrating Mutex Lock")
	resourceMutex := sync.Mutex{}

	go func() {
		resourceMutex.Lock()
		fmt.Println("Goroutine 1: Locked the resource")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1: Unlocking the resource")
		resourceMutex.Unlock()
	}()

	time.Sleep(500 * time.Millisecond)

	go func() {
		fmt.Println("Goroutine 2: Attempting to lock the resource")
		resourceMutex.Lock()
		defer resourceMutex.Unlock()
		fmt.Println("Goroutine 2: Locked the resource")
	}()

	time.Sleep(3 * time.Second)
}
