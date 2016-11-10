package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour",
	}
	fmt.Println(myGreeting["Jenny"])

	myGreeting["Jenny"] = "God morgon"

	fmt.Println(myGreeting["Jenny"])
}
