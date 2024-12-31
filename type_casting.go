package main

import (
	"Basics/utils"
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	utils.PrintSection("1. Numeric Type Casting")
	var intNum = 42
	var floatNum = float64(intNum)
	fmt.Printf("Int to Float: %v (type: %T)\n", floatNum, floatNum)

	var largeFloat = 1234.56
	var intFromFloat = int(largeFloat)
	fmt.Printf("Float to Int (truncated): %v (type: %T)\n", intFromFloat, intFromFloat)

	utils.PrintSection("2. String and Byte Slice Conversion")
	str := "Hello, Go!"
	byteSlice := []byte(str)
	fmt.Printf("String to Byte Slice: %v\n", byteSlice)

	newStr := string(byteSlice)
	fmt.Printf("Byte Slice to String: %s\n", newStr)

	utils.PrintSection("3. Interface Type Assertion")
	var i interface{} = "42"
	strVal, ok := i.(string)
	if ok {
		fmt.Printf("Interface to String: %s\n", strVal)
	}

	utils.PrintSection("4. Struct Type Conversion")
	type Rectangle struct {
		Width  float64
		Height float64
	}

	type Square struct {
		SideLength float64
	}

	square := Square{SideLength: 5}
	rectangle := Rectangle{
		Width:  square.SideLength,
		Height: square.SideLength,
	}

	fmt.Printf("Square: %+v\n", square)
	fmt.Printf("Rectangle (after conversion): %+v\n", rectangle)

	squareArea := square.SideLength * square.SideLength
	rectangleArea := rectangle.Width * rectangle.Height

	fmt.Printf("Square Area: %.2f\n", squareArea)
	fmt.Printf("Rectangle Area: %.2f\n", rectangleArea)

	utils.PrintSection("5. Pointer Type Conversion")
	x := int64(10)
	var p = &x
	var p2 = (*int32)(unsafe.Pointer(p))
	x = 12
	fmt.Printf("Original value (x): %v\n", x)
	fmt.Printf("Through p (*int64): %v\n", *p)
	fmt.Printf("Through p2 (*int32): %v\n", *p2)

	utils.PrintSection("6. String to Int Conversion")
	strNumber := "123"
	intNumber, err := strconv.Atoi(strNumber)
	if err == nil {
		fmt.Printf("String to Int: %d\n", intNumber)
	}

	utils.PrintSection("7. Int to String Conversion")
	num := 456
	strNum := strconv.Itoa(num)
	fmt.Printf("Int to String: %s\n", strNum)
}
