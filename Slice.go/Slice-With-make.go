package main

import (
	"fmt"
)

func main() {
	ms1 := make([]string, 5, 10)

	fmt.Printf("value of make slice: %v\n", ms1)
	fmt.Printf("length: %v\n", len(ms1))
	fmt.Printf("length: %v\n", cap(ms1))

}
