package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("wassup Julian / Sushant")
	}
}
