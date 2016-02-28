package main

import "fmt"

func main() {
	age := 344
	fmt.Println(&age) // 0xc082002298

	changeMe(&age)

	fmt.Println(&age) // 0xc082002298
	fmt.Println(age)  // 44
}

func changeMe(z *int) {
	fmt.Println(z)  // 0xc082002298
	fmt.Println(*z) // 24
	*z = 24
	fmt.Println(z)  // 0xc082002298
	fmt.Println(*z) // 24
}
