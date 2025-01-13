package main

import (
	"Basics/utils"
	"bytes"
	"fmt"
)

func main() {
	utils.PrintSection("1. Creating and Writing to Buffer")

	// Create a new buffer
	var buf bytes.Buffer

	// Write strings to buffer
	buf.WriteString("Hello, ")
	buf.WriteString("World!")

	// Display the buffer contents
	fmt.Printf("Buffer contents: %s\n", buf.String())

	utils.PrintSection("2. Writing Bytes and Slices")

	// Write a single byte
	buf.WriteByte(' ')

	// Write a byte slice
	buf.Write([]byte("From Go"))

	// Display the updated buffer contents
	fmt.Printf("Updated buffer contents: %s\n", buf.String())

	utils.PrintSection("3. Reading from Buffer")

	// Create a byte slice to read into
	readBytes := make([]byte, 5)
	n, err := buf.Read(readBytes)
	if err == nil {
		fmt.Printf("Read %d bytes: %s\n", n, string(readBytes))
	}

	// Read the next byte
	nextByte, err := buf.ReadByte()
	if err == nil {
		fmt.Printf("Next byte: %c\n", nextByte)
	}

	utils.PrintSection("4. Resetting and Reusing Buffer")

	// Reset the buffer
	buf.Reset()
	fmt.Printf("Buffer after reset: %q\n", buf.String())

	// Reuse buffer by writing new data
	buf.WriteString("Reused Buffer")
	fmt.Printf("Buffer contents after reuse: %s\n", buf.String())

	utils.PrintSection("5. Using Buffer for String Building")

	// Efficiently build a string using a buffer
	var stringBuf bytes.Buffer
	words := []string{"This", "is", "a", "buffered", "string"}
	for i, word := range words {
		if i > 0 {
			stringBuf.WriteByte(' ')
		}
		stringBuf.WriteString(word)
	}
	fmt.Printf("Built string: %s\n", stringBuf.String())
}
