package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "God morgon",
		2: "Hej",
		3: "Tjo",
	}

	for key, val := range myGreeting {
		fmt.Println(key, "-", val)
	}
}
