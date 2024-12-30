package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 1. Traditional for loop
	fmt.Println("1. Traditional for loop:")
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
		fmt.Printf("Iteration %d, Sum: %d\n", i, sum)
	}

	// 2. For loop as a while loop
	fmt.Println("\n2. For loop as a while loop:")
	num := 10
	for num > 0 {
		fmt.Printf("Countdown: %d\n", num)
		num--
	}

	// 3. Infinite loop with break
	fmt.Println("\n3. Infinite loop with break:")
	count := 0
	for {
		if count == 5 {
			break
		}
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// 4. For loop with continue
	fmt.Println("\n4. For loop with continue:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Odd number: %d\n", i)
	}

	// 5. For-range loop over a slice
	fmt.Println("\n5. For-range loop over a slice:")
	fruits := []string{"apple", "banana", "cherry", "date"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// 6. For-range loop over a map
	fmt.Println("\n6. For-range loop over a map:")
	scores := map[string]int{"Alice": 95, "Bob": 82, "Charlie": 88}
	for name, score := range scores {
		fmt.Printf("%s scored %d\n", name, score)
	}

	// 7. Nested loops
	fmt.Println("\n7. Nested loops:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	// 8. Loop with multiple variables
	fmt.Println("\n8. Loop with multiple variables:")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}

	// 9. For loop with a random number of iterations
	fmt.Println("\n9. For loop with a random number of iterations:")
	target := rand.Intn(10) + 1
	for attempts := 1; ; attempts++ {
		guess := rand.Intn(10) + 1
		if guess == target {
			fmt.Printf("Target %d found in %d attempts\n", target, attempts)
			break
		}
	}
}
