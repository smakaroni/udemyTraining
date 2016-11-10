package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	test := myGreeting["Jenny"]
	if test == "" {
		fmt.Println("NIL")
	} else {
		fmt.Println(test)
	}
	fmt.Println(myGreeting["Tims"])
}
