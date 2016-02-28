package main

import "fmt"

func main() {
	var largeNumber, smallnumber int
	fmt.Print("print a large number: ")
	fmt.Scan(&largeNumber)
	fmt.Print("print a small number: ")
	fmt.Scan(&smallnumber)
	fmt.Println(largeNumber / smallnumber)
}
