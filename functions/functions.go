package main

import (
	"errors"
	"fmt"
)

// Simple function with two parameters and one return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Closure example
func makeMultiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

// Anonymous function example
var subtract = func(a int, b int) int {
	return a - b
}

func main() {
	// Calling a simple function
	fmt.Println("Addition:", add(2, 3))

	// Calling a function with multiple returns
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}

	// Calling a variadic function
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5, 6, 7))

	// Using a closure
	multiplier := makeMultiplier(3)
	fmt.Println("Multiply:", multiplier(4))

	// Using an anonymous function
	fmt.Println("Subtraction:", subtract(10, 4))
}
