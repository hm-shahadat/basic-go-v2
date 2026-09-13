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

// ---------------------------

// short example

// package main

// import (
// 	"fmt"
// )

// func n(x *int) {
// 	*x = 35
// 	fmt.Println("it's not main function", *x)
// }

// func main() {
// 	x := 1

// 	n(&x)

// 	fmt.Println("it's main function", x)
// }
