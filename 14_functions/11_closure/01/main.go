package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Bla bla bla"
		fmt.Println(y)
	}
	// fmt.Println(y) outside of scope
}
