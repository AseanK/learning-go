package main

import "fmt"

// Key value pair
// both keys and values are statically typed
func main() {
	colors := map[int]string{
		000: "#ff0000",
	}
	fmt.Println(colors)
	// get specific value with key
	fmt.Println(colors[000])

	age := make(map[string]int)
	age["John"] = 32
	fmt.Println(age)

	// delete
	delete(colors, 000)
	fmt.Println(colors)
}
