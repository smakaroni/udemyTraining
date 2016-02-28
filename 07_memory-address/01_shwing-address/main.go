package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's momery address - ", &a)
	fmt.Printf("%d \n", &a)
	fmt.Printf("%b \n", &a)
}
