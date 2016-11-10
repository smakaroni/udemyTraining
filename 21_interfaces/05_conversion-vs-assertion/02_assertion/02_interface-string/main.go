package main

import "fmt"

func main() {
	var name interface{} = "Sidney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("Value not a string")
	}
}
