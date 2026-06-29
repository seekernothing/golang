package main

import "fmt"

func main() {
	/*
		age := 16

		if age >= 18 {
			fmt.Println("Person is adult")
		} else {
			fmt.Println("person is not an adult bro ")
		}

	*/

	/*
		age := 10

		if age >= 18 {
			fmt.Println("Person is an adult bro")
		} else if age >= 12 {
			fmt.Println("Person is a teenager")
		} else {
			fmt.Println("Person is a kid ")
		}

	*/

	role := "user"

	hasPermission := false

	if role == "admin" && hasPermission {
		fmt.Println("Yes")
	} else {
		fmt.Println("You are not allowed to check this route bri")
	}

	// go does not have ternary operator like javascript
}
