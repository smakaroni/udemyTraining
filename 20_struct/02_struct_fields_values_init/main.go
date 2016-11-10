package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"Jokke", "Persson", 28}
	p2 := person{"Fredrik", "Persson", 25}
	var p3 person // declaring a person with its zero value, idiomatic way to declare something to its zero value use var
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p3.first, p3.last, p3.age)
}
