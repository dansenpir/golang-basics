package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(a, b int) (int, error) {
	if b == 0 {
		// Returns an error if division by zero is detected
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Example of error usage
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example of panic and recover usage
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Function that causes a panic
	causePanic()
	fmt.Println("This line will not be executed if panic occurs")
}

func causePanic() {
	fmt.Println("About to panic")
	panic("something went wrong")
}
