package main

import (
	"fmt"
)

func main() {

	s1 := []int{}
	s2 := []string{"This", "is", "slice"}

	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)

	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2)

}
