package main

import "fmt"

func id() func() int {

	var count int = 0

	return func() int {

		count += 1
		return count
	}
}

func main() {

	incriment := id()

	fmt.Println(incriment())

}
