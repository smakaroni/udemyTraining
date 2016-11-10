package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("I'm a secret agent")
}

func main() {
	p1 := Person{
		Name: "Jokke Persson",
		Age: 28,
	}
	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age: 85,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
