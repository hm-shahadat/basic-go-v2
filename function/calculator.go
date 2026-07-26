package main

import (
	"fmt"
)

func add(x, y float32) float32 {
	return x + y
}

func sub(x, y float32) float32 {
	return x - y
}

func mul(x, y float32) float32 {
	return x * y
}

func div(x, y float32) float32 {
	return x / y
}

func main() {

	var n1, n2, result float32
	var operator string
	i := true

	for i == true {

		fmt.Printf("please an integer number:")
		fmt.Scan(&n1)

		fmt.Printf("please enter another integer number:")
		fmt.Scan(&n2)

		fmt.Printf("Please enter a operator (+ _ * /):")
		fmt.Scan(&operator)

		switch operator {
		case "+":
			result = add(n1, n2)

		case "-":
			result = sub(n1, n2)

		case "*":
			result = mul(n1, n2)

		case "/":
			result = div(n1, n2)
		default:
			fmt.Println("Invalid")
			continue
		}

		fmt.Printf("Result:%v\n", result)
	}

}
