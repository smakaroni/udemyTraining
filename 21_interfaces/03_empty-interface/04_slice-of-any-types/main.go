package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"Wof"}, true}
	fifi := cat{animal{"Meow"}, true}
	critters := []interface{}{fido, fifi}

	fmt.Println(critters)
}