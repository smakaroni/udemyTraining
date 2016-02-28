package main

import "fmt"

func main() {
	age := "Jokke"
	fmt.Println(&age) // 0xc082002298

	changeMe(&age)

	fmt.Println(&age) // 0xc082002298
	fmt.Println(age)  // 44
}

func changeMe(z *string) {
	fmt.Println(z)  // 0xc082002298
	fmt.Println(*z) // 24
	*z = "Chrystel"
	fmt.Println(z)  // 0xc082002298
	fmt.Println(*z) // 24
}
