package main

import (
	"Basics/utils"
	"fmt"
)

// Shape Interface type
type Shape interface {
	Area() float64
}

// Circle Struct types implementing Shape interface
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// Area Method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Area Method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	utils.PrintSection("1. Basic Types")

	var i int = 42
	var f float64 = 3.14
	var b bool = true
	var s string = "hello"

	fmt.Printf("int: %v, float64: %v, bool: %v, string: %v\n", i, f, b, s)

	utils.PrintSection("2. Aggregate Types (Array and Struct)")

	// Array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Struct:", p)

	utils.PrintSection("3. Reference Types (Slice, Map, and Pointer)")

	// Slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	// Map
	m := map[string]int{"one": 1, "two": 2}
	fmt.Println("Map:", m)

	// Pointer
	ptr := &i
	fmt.Println("Pointer:", ptr, "Value:", *ptr)

	utils.PrintSection("4. Interface Types and Type Assertions")

	var shape Shape
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 3, Height: 4}

	// Using interface
	shape = circle
	fmt.Printf("Circle area: %.2f\n", shape.Area())

	shape = rectangle
	fmt.Printf("Rectangle area: %.2f\n", shape.Area())

	// Type assertions
	checkShape(circle)
	checkShape(rectangle)
}

func checkShape(shape Shape) {

	fmt.Printf("Shape type: %T\n", shape)
	if circle, ok := shape.(Circle); ok {
		fmt.Printf("This is a circle with radius %.2f\n", circle.Radius)
	} else if rectangle, ok := shape.(Rectangle); ok {
		fmt.Printf("This is a rectangle with dimensions %.2f x %.2f\n", rectangle.Width, rectangle.Height)
	} else {
		fmt.Println("Unknown shape type")
	}
}
