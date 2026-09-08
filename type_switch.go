package main

import "fmt"

func main() {
	any := func(i interface{}) {

		switch t := i.(type) {
		case int:

			fmt.Println("integer")

		case bool:

			fmt.Println("boolean")

		case string:
			fmt.Println("string")

		default:
			fmt.Println("other", t)

		}
	}

	any(40.5)
}
