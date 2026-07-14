package main

import (
	"fmt"
)

func main() {

	var sc int
	fmt.Println("please enter your class")
	fmt.Scan(&sc)

	switch sc {

	case 1, 2, 3, 4, 5:

		fmt.Println("primary\n")

	case 6, 7, 8, 9, 10:

		fmt.Println("secondary\n")

	case 11, 12:

		fmt.Println("higher secondary\n")

	}

}
