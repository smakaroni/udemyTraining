package main

import "fmt"

func main() {

	// func expression
	greeting := func() {
		fmt.Println("Hello world")
	}
	greeting()
	fmt.Printf("%T\n", greeting)
}
