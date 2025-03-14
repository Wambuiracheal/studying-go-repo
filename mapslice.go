package main

import (
	"fmt"
)

func main() {
	// ====== SLICES ======
	fmt.Println("Slices in Go:")

	// Declare and initialize a slice
	numbers := []int{10, 20, 30, 40}
	fmt.Println("Slice: ", numbers)

	// Append elements to a slice
	numbers = append(numbers, 50, 60)
	fmt.Println("After appending: ", numbers)

	// Slice a slice
	subSlice := numbers[1:4] // From index 1 to 3
	fmt.Println("Sub-slice (1:4): ", subSlice)

	// Using make to create a slice with initial capacity
	newSlice := make([]int, 3, 5) // Length 3, Capacity 5
	fmt.Println("New Slice with make: ", newSlice, "Length:", len(newSlice), "Capacity:", cap(newSlice))

	// Modify elements
	numbers[2] = 99
	fmt.Println("After modification: ", numbers)

	// ====== MAPS ======
	fmt.Println("\nMaps in Go:")

	// Declare and initialize a map
	person := map[string]string{
		"name":    "Alice",
		"country": "USA",
	}
	fmt.Println("Map: ", person)

	// Add or update a key-value pair
	person["age"] = "25"
	fmt.Println("After adding age: ", person)

	// Retrieve value
	name := person["name"]
	fmt.Println("Name from map: ", name)

	// Delete a key
	delete(person, "country")
	fmt.Println("After deleting country: ", person)

	// Check if a key exists
	val, exists := person["age"]
	if exists {
		fmt.Println("Age exists with value: ", val)
	} else {
		fmt.Println("Age does not exist")
	}

	// Iterating over a map
	fmt.Println("\nIterating over map:")
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
