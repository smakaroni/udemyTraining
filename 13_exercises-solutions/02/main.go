package main

import "fmt"

func main() {
	var yourName string
	fmt.Print("enter your name: ")
	fmt.Scan(&yourName)
	fmt.Println("Hello ", yourName)
}
