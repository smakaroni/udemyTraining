package main

import "fmt"

func main() {

	// This only sets the capacity to the length
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}
