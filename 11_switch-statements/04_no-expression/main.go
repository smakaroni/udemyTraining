package main

import "fmt"

func main() {
	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	default:
		fmt.Println("Nothing matched; this is the default")

	}
}
