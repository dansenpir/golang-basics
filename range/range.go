package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Range over a slice of integers (original example)
	fmt.Println("1. Powers of 2:")
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	// 2. Range over a string
	fmt.Println("\n2. Characters in a string:")
	for i, char := range "Hello" {
		fmt.Printf("index: %d, character: %c\n", i, char)
	}

	// 3. Range with only index
	fmt.Println("\n3. Counting up:")
	for i := range 5 { // This is equivalent to: for i := 0; i < 5; i++
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 4. Range over array
	fmt.Println("\n4. Days of the week:")
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for _, day := range days {
		fmt.Println(day)
	}

	// 5. Using range to find elements
	fmt.Println("\n5. Finding vowels:")
	text := "Hello, World!"
	vowels := "aeiouAEIOU"
	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			fmt.Printf("%c is a vowel\n", char)
		}
	}

	// 6. Range to calculate sum
	fmt.Println("\n6. Sum of numbers:")
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Sum: %d\n", sum)

	// 7. Range with break
	fmt.Println("\n7. Finding the first even number:")
	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("First even number is at index %d: %d\n", i, num)
			break
		}
	}
}
