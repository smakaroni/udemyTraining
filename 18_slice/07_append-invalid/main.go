package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length
	// 5 is capacity
	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "dias"
	greeting[3] = "God morgon" // index out of range

	fmt.Println(greeting[2])
}
