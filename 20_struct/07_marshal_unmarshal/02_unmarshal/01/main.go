package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"Jokke", "Last":"Persson", "Age":20}`)
	fmt.Println(bs)
	json.Unmarshal(bs, &p1)

	fmt.Println("------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}