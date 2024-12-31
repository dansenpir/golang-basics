package main

import (
	"fmt"
	"strings"
)

func printSection(title string) {
	fmt.Printf("\n%s\n%s\n", title, strings.Repeat("-", len(title)))
}

func main() {
	printSection("1. Declaring and Initializing Arrays")
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Array of numbers:", numbers)

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Array of fruits:", fruits)

	printSection("2. Accessing Array Elements")
	fmt.Println("First number:", numbers[0])
	fmt.Println("Second fruit:", fruits[1])

	printSection("3. Iterating Over Arrays")
	fmt.Println("Numbers array:")
	for i, v := range numbers {
		fmt.Printf("  Index: %d, Value: %d\n", i, v)
	}

	printSection("4. Array Length")
	fmt.Println("Length of numbers array:", len(numbers))
	fmt.Println("Length of fruits array:", len(fruits))

	printSection("5. Multidimensional Arrays")
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2x3 Matrix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("  %d ", matrix[i][j])
		}
		fmt.Println()
	}

	printSection("6. Array Comparison")
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{3, 2, 1}
	fmt.Printf("arr1 == arr2: %v\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %v\n", arr1 == arr3)

	printSection("7. Copying Arrays")
	arrOriginal := [3]int{1, 2, 3}
	arrCopy := arrOriginal
	arrCopy[0] = 100
	fmt.Println("Original array:", arrOriginal)
	fmt.Println("Copied and modified array:", arrCopy)
}
