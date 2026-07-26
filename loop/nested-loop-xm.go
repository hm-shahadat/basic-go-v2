package main

import "fmt"

func main() {

	var a int
	fmt.Println("please enter n-th integer number:")
	fmt.Scan(&a)

	for i := a; i >= 1; i-- { i=5, 

		for j := 1; j <= i; j++ { j = 1 
			fmt.Print("*")
		}

		fmt.Println()
	}
}
