package main

import "fmt"

func main() {
	n := (true && false) || (false && true) || !(false && false)
	fmt.Println(n)
}
