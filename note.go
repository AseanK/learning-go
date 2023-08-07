package main

import "fmt"

// String, Int, Float, Bool, Struts
// Pointers are pointing at the value

// Slices, Maps, Channels, Functions
// Pointers are pointing at the RAM
func main() {
	n := "Bill"

	update(&n)
	fmt.Println(n)
}

func update(n *string) {
	*n = "Ill"
}
