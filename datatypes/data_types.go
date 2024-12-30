package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

var (
	// Boolean
	toBe bool = false

	// Signed integers
	maxInt8  int8  = math.MaxInt8
	maxInt16 int16 = math.MaxInt16
	maxInt32 int32 = math.MaxInt32
	maxInt64 int64 = math.MaxInt64
	maxInt   int   = int(^uint(0) >> 1)

	// Unsigned integers
	maxUint8  uint8  = math.MaxUint8
	maxUint16 uint16 = math.MaxUint16
	maxUint32 uint32 = math.MaxUint32
	maxUint64 uint64 = math.MaxUint64
	maxUint   uint   = ^uint(0)

	// Floating point
	maxFloat32 float32 = math.MaxFloat32
	maxFloat64 float64 = math.MaxFloat64

	// Complex numbers
	complexNum64  complex64  = complex(1, 2)
	complexNum128 complex128 = cmplx.Sqrt(-5 + 12i)

	// String
	hello string = "Hello, Gopher!"

	// Byte (alias for uint8) and Rune (alias for int32)
	byteVal byte = 255
	runeVal rune = '🐹'

	// Pointer
	uintptrVal uintptr = uintptr(unsafe.Pointer(&maxInt))
)

func main() {
	fmt.Println("Boolean Types:")
	fmt.Printf("bool: \t\tType: %T \tValue: %v\n\n", toBe, toBe)

	fmt.Println("Signed Integer Types:")
	fmt.Printf("int8: \t\tType: %T \tValue: %v\n", maxInt8, maxInt8)
	fmt.Printf("int16: \t\tType: %T \tValue: %v\n", maxInt16, maxInt16)
	fmt.Printf("int32: \t\tType: %T \tValue: %v\n", maxInt32, maxInt32)
	fmt.Printf("int64: \t\tType: %T \tValue: %v\n", maxInt64, maxInt64)
	fmt.Printf("int: \t\tType: %T \tValue: %v\n\n", maxInt, maxInt)

	fmt.Println("Unsigned Integer Types:")
	fmt.Printf("uint8: \t\tType: %T \tValue: %v\n", maxUint8, maxUint8)
	fmt.Printf("uint16: \tType: %T \tValue: %v\n", maxUint16, maxUint16)
	fmt.Printf("uint32: \tType: %T \tValue: %v\n", maxUint32, maxUint32)
	fmt.Printf("uint64: \tType: %T \tValue: %v\n", maxUint64, maxUint64)
	fmt.Printf("uint: \t\tType: %T \tValue: %v\n\n", maxUint, maxUint)

	fmt.Println("Floating Point Types:")
	fmt.Printf("float32: \tType: %T \tValue: %v\n", maxFloat32, maxFloat32)
	fmt.Printf("float64: \tType: %T \tValue: %v\n\n", maxFloat64, maxFloat64)

	fmt.Println("Complex Number Types:")
	fmt.Printf("complex64: \tType: %T \tValue: %v\n", complexNum64, complexNum64)
	fmt.Printf("complex128: \tType: %T \tValue: %v\n\n", complexNum128, complexNum128)

	fmt.Println("String Type:")
	fmt.Printf("string: \tType: %T \tValue: %v\n\n", hello, hello)

	fmt.Println("Alias Types:")
	fmt.Printf("byte: \t\tType: %T \tValue: %v\n", byteVal, byteVal)
	fmt.Printf("rune: \t\tType: %T \tValue: %v\n\n", runeVal, runeVal)

	fmt.Println("Pointer Type:")
	fmt.Printf("uintptr: \tType: %T \tValue: %v\n", uintptrVal, uintptrVal)
}
