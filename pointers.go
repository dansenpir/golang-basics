package main

import (
	"fmt"
	"strings"
)

// Individual represents a person with a name and age
type Individual struct {
	Name string
	Age  int
}

// UpdateAge increases the individual's age by one year
func (i *Individual) UpdateAge() {
	i.Age++
}

// UpdateName changes the individual’s name to uppercase
func UpdateName(i *Individual) {
	i.Name = strings.ToUpper(i.Name)
}

// SwapAges swaps the ages of two individuals
func SwapAges(i1, i2 *Individual) {
	i1.Age, i2.Age = i2.Age, i1.Age
}

// PrintIndividualInfo prints information about an individual
func PrintIndividualInfo(i *Individual) {
	fmt.Printf("Name: %s, Age: %d\n", i.Name, i.Age)
}

// LargeDataStruct represents a structure with a large amount of data
type LargeDataStruct struct {
	Data [1000000]int
}

// ModifyLargeData modifies the first element of the large data structure
func ModifyLargeData(lds *LargeDataStruct) {
	lds.Data[0] = 100
}

func main() {
	// Basic pointer usage
	x := 42
	ptr := &x
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", ptr)
	fmt.Printf("Value at address stored in ptr: %d\n", *ptr)

	*ptr = 100
	fmt.Printf("New value of x: %d\n", x)

	// Using pointers with structs
	alice := Individual{Name: "Alice", Age: 30}
	bob := Individual{Name: "Bob", Age: 25}

	fmt.Println("\nBefore updates:")
	PrintIndividualInfo(&alice)
	PrintIndividualInfo(&bob)

	alice.UpdateAge()
	UpdateName(&bob)
	SwapAges(&alice, &bob)

	fmt.Println("\nAfter updates:")
	PrintIndividualInfo(&alice)
	PrintIndividualInfo(&bob)

	// Using pointers for efficiency with large data structures
	largeData := LargeDataStruct{}
	fmt.Printf("\nBefore modification: %d\n", largeData.Data[0])
	ModifyLargeData(&largeData)
	fmt.Printf("After modification: %d\n", largeData.Data[0])

	// Slices and pointers
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("\nOriginal slice: %v\n", numbers)

	// Modifying slice without using a pointer
	modifySlice := func(s []int) {
		s[0] = 100
	}
	modifySlice(numbers)
	fmt.Printf("After modifying slice: %v\n", numbers)
}
