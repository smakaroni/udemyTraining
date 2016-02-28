package main

import "fmt"

func main() {
	n := avarage(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

// ... means its variadic
func avarage(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
