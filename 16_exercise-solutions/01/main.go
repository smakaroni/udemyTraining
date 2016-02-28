package main

import "fmt"

func main() {
	fmt.Println(devideIt(14))
}

func devideIt(z int) (int, bool) {
	return z / 2, z%2 == 0

}
