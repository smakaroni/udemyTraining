package main

import (
	"strconv"
	"fmt"
)

func main() {
	x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y)
}
