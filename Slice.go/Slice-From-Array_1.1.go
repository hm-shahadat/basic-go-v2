package main

import (
	"fmt"
)

func main() {
	//ARRAY
	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a1)

	//-----Create a Slice From an Array------
	a2 := [5]int{10, 20, 30, 40, 50}
	s1 := a2[1:3]

	// fmt.Println(a2)
	// fmt.Println(s1)

	fmt.Printf("slice value: %v\n", s1)
	fmt.Printf("slice type: %T\n", s1)
	fmt.Printf("slice length: %d\n", len(s1))
	fmt.Printf("slice capacity: %d\n", cap(s1))

}
