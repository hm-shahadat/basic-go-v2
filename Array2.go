package main

import "fmt"

func main() {

	var arr [3]string

	// Print Array using loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Print("Please enter 5 student name: ")
	// 	fmt.Scan(&arr[i])

	// }

	// fmt.Println(arr)

	// user input index number and then print that index value
	// var choice int
	// fmt.Print("please enter e number: ")
	// fmt.Scan(&choice)

	// fmt.Println(arr[choice])

	for i := 0; i < 3; i++ {

		fmt.Print("please enter your full name in 3 step:")
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
}
