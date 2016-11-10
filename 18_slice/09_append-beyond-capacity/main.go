package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length
	// 5 is capacity
	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "dias"
	greeting = append(greeting, "God morgon")
	greeting = append(greeting, "Gidday")
	greeting = append(greeting, "Morn")
	greeting = append(greeting, "suprabadham")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
