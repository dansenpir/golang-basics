package main

import (
	"Basics/utils"
	"fmt"
)

// PrintSlice Generic function to print any slice
func PrintSlice[T any](slice []T) {
	fmt.Print("[ ")
	for _, v := range slice {
		fmt.Printf("%v ", v)
	}
	fmt.Println("]")
}

// FindMax Generic function to find the maximum value in a slice
func FindMax[T ~int | ~float64 | ~string](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	maxVal := slice[0]
	for _, v := range slice[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal, true
}

// Stack Generic Stack data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Number Generic function with constraints
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	utils.PrintSection("1. Generic PrintSlice Function")
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"}
	PrintSlice(intSlice)
	PrintSlice(stringSlice)

	utils.PrintSection("2. Generic FindMax Function")
	if maxInt, ok := FindMax(intSlice); ok {
		fmt.Printf("Max of int slice: %d\n", maxInt)
	}
	if maxString, ok := FindMax(stringSlice); ok {
		fmt.Printf("Max of string slice: %s\n", maxString)
	}

	utils.PrintSection("3. Generic Stack")
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	for {
		if item, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\n", item)
		} else {
			break
		}
	}

	utils.PrintSection("4. Generic Sum Function with Constraints")
	intNumbers := []int{1, 2, 3, 4, 5}
	floatNumbers := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Printf("Sum of ints: %d\n", Sum(intNumbers))
	fmt.Printf("Sum of floats: %.2f\n", Sum(floatNumbers))
}
