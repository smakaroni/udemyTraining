package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Jokke"])
}

func changeMe(z map[string]int) {
	z["Jokke"] = 27
}
