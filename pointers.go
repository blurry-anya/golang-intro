package main

import "fmt"

func setName(n *string) {
	*n = "Abigail"
}

func pointers() {
	name := "Anya"
	fmt.Println("Original name:", name)

	fmt.Println("Memory address of name:", &name)

	m := &name
	fmt.Println("Memory address of m:", m) // pointer
	fmt.Println("Value stored in m:", *m)  // reference value
	setName(m)                             // basically setName(&name)

	fmt.Println("Updated name:", name)

}
