package main

import (
	"fmt"
)

func main() {

	// var array1 = [2]int{10, 20}
	a := [...]string{"hello", "world", "universe", "galaxy"}
	// b := [5]int{1: 10, 2: 20}
	// a[1] = "galaxy"
	// fmt.Printf("the value: %v\nthe type: %T\n", array2, array2)
	// fmt.Printf("the value: %v\nthe type: %T\n", array1, array1)

	// fmt.Println(a[2])
	// fmt.Println(a[1])
	// fmt.Println(b[1])
	fmt.Println(len(a))
}
