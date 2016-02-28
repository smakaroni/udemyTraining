package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 42 // b says, "The value at this address, change it to 42"
	fmt.Println(a)
}
