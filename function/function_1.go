package main

import (
	"fmt"
)

func sq(num1 int, num2 int) int {
	result := num1 * num2

	return result
}

func sub(n1 int, n2 int) {

	result := n1 - n2
	fmt.Printf("%v\n", result)
}

func n() string {

	return "with return"

}
func main() {

	fmt.Printf("%v\n", sq(2, 2))
	sub(5, 3)
	fmt.Println(n())
}
