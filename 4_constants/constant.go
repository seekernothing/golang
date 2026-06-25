package main

import "fmt"

// 2 = you can declare constants outside of main function

const salary = 777

func main() {

	/*

		3 =   "  :=  " you can't use this shorthand outside of function

		                & in const"

	*/

	/*
		1 = 	const name = "Golang"

			name = "Javascript"

	*/

	//age := 20

	/*
		     4 = if you want to use multiple constant

			 you can use conatnt grouping


	*/

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(salary)

	fmt.Println("Port is", port, "host is", host)

}
