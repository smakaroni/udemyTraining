package main

import "fmt"

func makeGreater() func() string {
	return func() string {
		return "Hello world"
	}
}

func main() {
	greet := makeGreater()
	fmt.Println(greet())
}
