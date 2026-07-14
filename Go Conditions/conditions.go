package main

import (
	"fmt"
)

func main() {

	// a program to determine positive / negative / zero
	// var pnz int
	// fmt.Printf("please enter a number:")
	// fmt.Scan(&pnz)

	// if pnz > 0 {
	// 	fmt.Println("the number is positive")

	// } else if pnz == 0 {

	// 	fmt.Println("the number is zero")

	// } else {
	// 	fmt.Println("the number is negative")
	// }

	//  even/odd
	// var eo int
	// fmt.Printf("please enter a number:")
	// fmt.Scanf("%d", &eo)

	// if eo%2 == 0 {

	// 	fmt.Printf("%v is even\n", eo)

	// } else {

	// 	fmt.Printf("%v is odd\n", eo)
	// }

	//  find largest number among 3 numbers
	var ln1 int
	var ln2 int
	var ln3 int
	fmt.Printf("please enter 3 number:\n")
	fmt.Scan(&ln1, &ln2, &ln3)

	if (ln1 > ln2) && (ln1 > ln3) {
		fmt.Printf("%v is the largest number\n", ln1)

	} else if (ln2 > ln1) && (ln2 > ln3) {
		fmt.Printf("%v is the largest number\n", ln2)

	} else {
		fmt.Printf("%v is the largest number\n", ln3)
	}

}
