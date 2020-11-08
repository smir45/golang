package main

import (
	"fmt"
)

func main() {
	//comment
	fmt.Println("Enter your First Name: ")

	// we should write var for declaring vaiables
	// and variable name and then it's type
	var firstname string

	//Now taking input from user

	fmt.Scanln(&firstname)
	fmt.Println("Enter your Last Name: ")

	var lastname string

	fmt.Scanln(&lastname)

	//Print function is used to display output

	fmt.Println("your full name is:- ")
	fmt.Println(firstname + " " + lastname)

}
