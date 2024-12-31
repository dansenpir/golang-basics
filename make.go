package main

import (
	"Basics/utils"
	"fmt"
)

func main() {
	utils.PrintSection("1. Using make with Slices")

	// Create a slice with make
	// Syntax: make([]T, length, capacity)
	slice1 := make([]int, 5)
	fmt.Println("Slice1 (length 5, capacity 5):", slice1)

	// Create a slice with length different from capacity
	slice2 := make([]int, 3, 7)
	fmt.Println("Slice2 (length 3, capacity 7):", slice2)
	fmt.Printf("Slice2 length: %d, capacity: %d\n", len(slice2), cap(slice2))

	// Adding elements to slice2
	slice2 = append(slice2, 1, 2, 3, 4)
	fmt.Println("Slice2 after append:", slice2)

	utils.PrintSection("2. Using make with Maps")

	// Create a map with make
	// Syntax: make(map[KeyType]ValueType)
	map1 := make(map[string]int)
	fmt.Println("Empty map1:", map1)

	// Add elements to map1
	map1["apple"] = 5
	map1["banana"] = 3
	fmt.Println("Map1 after adding elements:", map1)

	// Create a map with initial size hint
	map2 := make(map[string]int, 10)
	fmt.Println("Empty map2 with size hint:", map2)
	// Size hint doesn't limit the map, it's just for initial memory allocation
	for i := 0; i < 15; i++ {
		map2[fmt.Sprintf("key%d", i)] = i
	}
	fmt.Println("Map2 after adding 15 elements:", map2)
}
