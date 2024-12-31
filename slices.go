package main

import (
	"Basics/utils"
	"fmt"
)

func main() {
	utils.PrintSection("1. Creating Slices")
	// Creating a slice using make
	numbers := make([]int, 5, 10)
	fmt.Printf("Numbers: %v\nLength: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

	// Creating a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := arr[1:4]
	fmt.Printf("Slice from array: %v\n", sliceFromArray)

	utils.PrintSection("2. Appending to Slices")
	fruits := []string{"Apple", "Banana"}
	fruits = append(fruits, "Cherry")
	fmt.Printf("Fruits after append: %v\n", fruits)

	utils.PrintSection("3. Slicing Slices")
	numberSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", numberSlice)
	fmt.Printf("First three: %v\n", numberSlice[:3])
	fmt.Printf("Last three: %v\n", numberSlice[len(numberSlice)-3:])
	fmt.Printf("Middle: %v\n", numberSlice[3:7])

	utils.PrintSection("4. Copying Slices")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Printf("Source: %v\nDestination: %v\nNumber of elements copied: %d\n", src, dst, copied)

	utils.PrintSection("5. Slice Capacity")
	s := make([]int, 3, 5)
	fmt.Printf("Initial: len=%d cap=%d %v\n", len(s), cap(s), s)
	s = append(s, 4)
	fmt.Printf("Append 4: len=%d cap=%d %v\n", len(s), cap(s), s)
	s = append(s, 5, 6, 7)
	fmt.Printf("Append 5,6,7: len=%d cap=%d %v\n", len(s), cap(s), s)

	utils.PrintSection("6. Nil Slices")
	var nilSlice []int
	fmt.Printf("Nil slice: %v, Is nil? %t\n", nilSlice, nilSlice == nil)

	utils.PrintSection("7. Using range with Slices")
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	utils.PrintSection("8. Multidimensional Slices")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Slice (Matrix):")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
