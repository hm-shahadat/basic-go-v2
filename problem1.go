package main

import "fmt"

func main() {

	x := 10

	p := &x
	*p = 15
	fmt.Println(*p)
	fmt.Println(p)

	// fmt.Println(x)
	// x++
	// fmt.Println(x)

}
