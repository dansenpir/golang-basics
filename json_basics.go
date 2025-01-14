package main

import (
	"encoding/json"
	"fmt"
)

// Book represents a book with title, author, and publication year
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// Library represents a collection of books
type Library struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func main() {
	// Create a library with some books
	myLibrary := Library{
		Name: "My Personal Library",
		Books: []Book{
			{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: 2015},
			{Title: "1984", Author: "George Orwell", Year: 1949},
		},
	}

	// Marshal: Convert Go struct to JSON
	jsonData, err := json.Marshal(myLibrary)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}
	fmt.Println("Marshalled JSON:")
	fmt.Println(string(jsonData))

	// Unmarshal: Convert JSON back to Go struct
	var newLibrary Library
	err = json.Unmarshal(jsonData, &newLibrary)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	fmt.Println("\nUnmarshalled Library:")
	fmt.Printf("Name: %s\n", newLibrary.Name)
	for _, book := range newLibrary.Books {
		fmt.Printf("- %s by %s (%d)\n", book.Title, book.Author, book.Year)
	}

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(myLibrary, "", "  ")
	if err != nil {
		fmt.Printf("Error creating pretty JSON: %v\n", err)
		return
	}
	fmt.Println("\nPretty Printed JSON:")
	fmt.Println(string(prettyJSON))
}
