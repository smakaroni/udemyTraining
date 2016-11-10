package main

import (
	"fmt"
	"encoding/json"
)
// values dont travel outside package(json is in another pkg) bcs this is unexported (lower case)
type Person struct {
	first string
	last string
	age int
}

func main() {
	p1 := Person{"Jokke", "Persson", 28}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
