package main

import "fmt"

func main() {
	var name interface{} = "Sidney"
	fmt.Println(name.(int) + 6)
}