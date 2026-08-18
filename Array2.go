package main

import "fmt"

func main() {

	var arr [5]string

	for i := 0; i < 5; i++ {
		fmt.Print("Please enter 5 student name: ")
		fmt.Scan(&arr[i])

	}

	fmt.Println(arr)

}
