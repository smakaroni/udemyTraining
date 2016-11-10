package main

import "fmt"

type Person struct {
	First string
	Last string
	Age int
}

type DoubleZero struct {
	Person
	First string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Jokke",
			Last: "Persson",
			Age: 28,
		},
		First: "007",
		LicenseToKill: true,
	}
	fmt.Println(p1.Person.First, p1.First, p1.Last, p1.Age)

}