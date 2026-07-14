package main

import (
	"fmt"
)

func main() {

	var a = 12 + 13
	fmt.Println(a)
	var (
		b = 2 + 5
		c = b + 2
		d = c + b
	)
	fmt.Println(d)

	// Arithmetic Operators
	var e = 2
	var f = 3
	var g = e % f
	fmt.Println(g)

	// Assignment Operators
	var h = 10
	h += 2
	fmt.Println(h)

	var i = 3
	i &= 2
	fmt.Println(i)

	// Comparison Operators
	j := 2
	k := 3
	fmt.Println(j < k)

	// Logical Operators
	var l = 10
	fmt.Println(l < 9 && l > 12)

	var x = 10
	fmt.Println(x < 11 || x > 12)

	var y = 10
	fmt.Println(!(y < 2 && y > 9))

	// Bitwise Operators
	var m = 9
	var n = 8
	fmt.Printf("m= %b\n", m)
	fmt.Printf("n= %b\n", n)
	fmt.Printf("m and n: %b\n", m&n)

}
