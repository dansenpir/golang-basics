package main

import (
	"Basics/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	utils.PrintSection("1. Basic Goroutine Usage")

	// Start a goroutine
	go sayHello("World")
	sayHello("Main")

	// Sleep to allow the goroutine to finish
	time.Sleep(100 * time.Millisecond)

	utils.PrintSection("2. Multiple Goroutines")

	for i := 0; i < 5; i++ {
		go printNumber(i)
	}
	time.Sleep(100 * time.Millisecond)

	utils.PrintSection("3. Goroutines with Channels")

	ch := make(chan string)
	go sendMessage(ch, "Hello from goroutine!")
	message := <-ch
	fmt.Println("Received:", message)

	utils.PrintSection("4. Goroutines with WaitGroup")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printNumberWithWaitGroup(i, &wg)
	}
	wg.Wait()

	utils.PrintSection("5. Goroutine with Closure")

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("Closure: %d\n", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func printNumber(n int) {
	fmt.Printf("Number: %d\n", n)
}

func sendMessage(ch chan string, message string) {
	ch <- message
}

func printNumberWithWaitGroup(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WaitGroup Number: %d\n", n)
}
