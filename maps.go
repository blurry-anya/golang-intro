package main

import "fmt"

func mapsDemo() {
	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   6.99,
		"pudding": 3.45,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, value := range menu {
		fmt.Printf("%v : %v\n", key, value)
	}

	// ints as key type
	phonebook := map[int]string{
		123456: "anya",
		654321: "abigail",
		125634: "robin",
		563412: "sebastian",
	}

	fmt.Println("\nPhonebook:")
	for key, value := range phonebook {
		fmt.Printf("%v : %v\n", key, value)
	}

	phonebook[563412] = "seba"

	fmt.Println(phonebook)

}
