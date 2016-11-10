package main

import "fmt"

func main() {
	var name interface{} = 7
	fmt.Printf("%T\n", name)
}