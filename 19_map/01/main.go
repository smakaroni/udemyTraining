package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)

	// Is the key in the map
	_, ok := m["k2"]
	fmt.Println("ok?:", ok)

	// one line decleration
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
