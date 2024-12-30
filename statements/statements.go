package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a new random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random temperature between -10 and 40
	temperature := r.Intn(51) - 10

	fmt.Printf("Current temperature: %d°C\n", temperature)

	// Using if-else statements
	if temperature < 0 {
		fmt.Println("It's freezing outside!")
	} else if temperature < 10 {
		fmt.Println("It's very cold. Wear a heavy coat.")
	} else if temperature < 20 {
		fmt.Println("It's cool. A light jacket should be fine.")
	} else if temperature < 30 {
		fmt.Println("It's a pleasant day. Enjoy the weather!")
	} else {
		fmt.Println("It's hot! Stay hydrated and find some shade.")
	}

	// Using multiple if statements for comprehensive analysis
	fmt.Println("\nDetailed weather analysis:")
	if temperature < 0 {
		fmt.Println("- Risk of snow or ice")
		fmt.Println("- Extreme cold weather precautions needed")
	}
	if temperature < 10 {
		fmt.Println("- Cold weather clothing recommended")
	}
	if temperature >= 10 && temperature < 20 {
		fmt.Println("- Cool weather, comfortable for most activities")
	}
	if temperature >= 20 && temperature < 30 {
		fmt.Println("- Warm weather, good for outdoor activities")
	}
	if temperature >= 30 {
		fmt.Println("- High temperature warning")
		fmt.Println("- Stay hydrated and avoid prolonged sun exposure")
	}

	// Using switch statement
	fmt.Println("\nActivity recommendations:")
	switch {
	case temperature < 0:
		fmt.Println("Winter activities: Ice skating or staying indoors.")
	case temperature < 15:
		fmt.Println("Outdoor activities: Brisk walking or jogging.")
	case temperature < 25:
		fmt.Println("Outdoor activities: Picnicking or playing sports.")
	default:
		fmt.Println("Outdoor activities: Swimming or staying in air-conditioned areas.")
	}

	// Switch statement with explicit values
	switch temperature {
	case 0:
		fmt.Println("\nTemperature is exactly at freezing point.")
	case 10:
		fmt.Println("\nTemperature is 10°C, a benchmark cool temperature.")
	case 20:
		fmt.Println("\nTemperature is 20°C, a benchmark comfortable temperature.")
	case 30:
		fmt.Println("\nTemperature is 30°C, a benchmark warm temperature.")
	}
}
