package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating a channel for integers
	messageChannel := make(chan int)

	// Goroutine that sends data
	go func() {
		fmt.Println("Goroutine: Starting to send numbers...")
		for i := 1; i <= 5; i++ {
			fmt.Printf("Goroutine: Sending %d\n", i)
			messageChannel <- i     // Sending a number to the channel
			time.Sleep(time.Second) // Waiting 1 second between each send
		}
		close(messageChannel) // Closing the channel after sending all numbers
	}()

	// Main goroutine that receives data
	fmt.Println("Main: Waiting for numbers...")
	for num := range messageChannel {
		fmt.Printf("Main: Received %d\n", num)
	}

	fmt.Println("Main: All numbers have been received. Program finished.")
}
