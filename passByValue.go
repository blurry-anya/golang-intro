package main

import "fmt"

func updateName(n string) string {
	n = "Abigail"
	return n
}

func updateMenu(m map[string]float64) {
	m["coffee"] = 2.5
}

func passByValue() {
	// Non-Pointer Values: types => strings, int, bools, floats, arrays, structs
	name := "Anya"
	fmt.Println("Original name:", name)

	name = updateName(name)

	fmt.Println("Updated name:", name)

	// Pointer Wrapper Values: types => maps, slices, functions
	menu := map[string]float64{
		"ice cream": 3.99,
		"pie":       2.99,
	}

	fmt.Println("Original menu:", menu)

	updateMenu(menu)

	fmt.Println("Updated menu:", menu)

}
