package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println(p.firstName, "wishes you a great day")
}

func (sa secretAgent) sSpeak() {
	fmt.Println(sa.firstName, sa.lastName, "has license to kill:", sa.licenseToKill)

}

func main() {
	p := person{"Jokke", "Persson"}
	sa := secretAgent{person{"James", "Bond"}, true}
	fmt.Println(p.firstName)
	// fmt.Println(sa.lastName)
	//p.pSpeak()
	sa.pSpeak()
	sa.sSpeak()
}
