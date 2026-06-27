package main

import "fmt"

func main() {

	// there is no while loop in go

	// you have to use for loop even for while loop

	// while loop

	/*i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}
	*/

	// classics for loop

	

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	

	for i := range 31 {

		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}




}
