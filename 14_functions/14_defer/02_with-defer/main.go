package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world")
}

// defer = senarelägg till precis innan main avslutas
func main() {
	defer world()
	hello()
}
