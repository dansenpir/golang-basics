package main

import (
	"Basics/utils"
	"fmt"
)

func main() {
	utils.PrintSection("1. Declaring and Initializing Maps")

	var emptyMap map[string]int
	fmt.Println("Empty map:", emptyMap)

	ages := make(map[string]int)
	fmt.Println("Ages (empty):", ages)

	populations := map[string]int{
		"New York":    8419000,
		"Los Angeles": 3898000,
		"Chicago":     2746000,
	}
	fmt.Println("Populations:", populations)

	utils.PrintSection("2. Adding and Updating Elements")

	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Println("Ages after adding:", ages)

	ages["Alice"] = 31 // Updating
	fmt.Println("Ages after updating Alice:", ages)

	utils.PrintSection("3. Accessing and Deleting Elements")

	aliceAge := ages["Alice"]
	fmt.Println("Alice's age:", aliceAge)

	bobAge, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age:", bobAge)
	} else {
		fmt.Println("Bob's age not found")
	}

	delete(ages, "Bob")
	fmt.Println("Ages after deleting Bob:", ages)

	utils.PrintSection("4. Iterating Over Maps")

	for city, population := range populations {
		fmt.Printf("%s: %d\n", city, population)
	}

	utils.PrintSection("5. Length of a Map")

	fmt.Println("Number of cities:", len(populations))

	utils.PrintSection("6. Maps as References")

	copyPopulations := populations
	copyPopulations["Seattle"] = 724000
	fmt.Println("Original populations (also modified):", populations)
	fmt.Println("Copy of populations:", copyPopulations)

	utils.PrintSection("7. Maps of Maps")

	cityInfo := map[string]map[string]string{
		"New York": {"country": "USA", "landmark": "Statue of Liberty"},
		"Paris":    {"country": "France", "landmark": "Eiffel Tower"},
	}
	fmt.Println("City Info:")
	for city, info := range cityInfo {
		fmt.Printf("%s - Country: %s, Landmark: %s\n", city, info["country"], info["landmark"])
	}
}
