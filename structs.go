package main

import (
	"Basics/utils"
	"fmt"
)

// Person Define a struct
type Person struct {
	Name string
	Age  int
	City string
}

// Introduce Method for Person struct
func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s, %d years old, from %s", p.Name, p.Age, p.City)
}

// Employee Another struct with embedded struct
type Employee struct {
	Person
	JobTitle string
	Salary   float64
}

func main() {
	utils.PrintSection("1. Creating and Using Structs")

	// Create a new Person
	p1 := Person{Name: "Alice", Age: 30, City: "New York"}
	fmt.Println("Person 1:", p1)

	// Create a Person with partial fields
	p2 := Person{Name: "Bob"}
	fmt.Println("Person 2:", p2)

	// Access and modify struct fields
	p2.Age = 25
	p2.City = "San Francisco"
	fmt.Println("Updated Person 2:", p2)

	utils.PrintSection("2. Struct Methods")

	// Call method on struct
	fmt.Println(p1.Introduce())

	utils.PrintSection("3. Nested Structs")

	// Create an Employee
	e1 := Employee{
		Person:   Person{Name: "Charlie", Age: 35, City: "Chicago"},
		JobTitle: "Developer",
		Salary:   75000,
	}
	fmt.Println("Employee:", e1)
	fmt.Println("Employee Name:", e1.Name) // Accessing embedded struct field

	utils.PrintSection("4. Pointers to Structs")

	// Create a pointer to a struct
	p3 := &Person{Name: "David", Age: 40, City: "Boston"}
	fmt.Println("Person Pointer:", p3)
	fmt.Println("Person Name via Pointer:", (*p3).Name)
	// Shorthand notation
	fmt.Println("Person Age via Pointer:", p3.Age)

	utils.PrintSection("5. Anonymous Structs")

	// Create an anonymous struct
	book := struct {
		Title  string
		Author string
		Pages  int
	}{
		Title:  "Go Programming",
		Author: "John Doe",
		Pages:  300,
	}
	fmt.Println("Book:", book)
}
