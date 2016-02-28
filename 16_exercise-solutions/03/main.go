package main

import "fmt"

func main() {
	n := findHighest(1, 2, 3, 4, 5)
	fmt.Println(n)
}

func findHighest(numbs ...int) int {
	var highest int
	for _, v := range numbs {
		if v > highest {
			highest = v
		}
	}
	return highest
}
