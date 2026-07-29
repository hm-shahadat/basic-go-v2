package main

import (
	"fmt"
)

func n(x int) {

	x = 35

}
func p(y *int) {
	*y = 60
}

func main() {

	x := 10
	n(x)
	fmt.Println(x)

	p(&x)
	fmt.Println(x)
}
