package main

import "fmt"

func main() {

	var a int
	fmt.Println("please enter n-th integer number:")
	fmt.Scan(&a)

	for i := a; i >= 1; i-- {

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
