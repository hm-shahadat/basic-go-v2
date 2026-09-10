package main

import "fmt"

func main() {

	//using slice
	slice := []int{6, 7, 8}

	for i, num := range slice {

		fmt.Println(num, i)
	}

	//using map

	m := map[string]string{"name": "shahadat", "position": "first"}

	for _, v := range m {

		fmt.Println(v)
	}

	// print both value and index name or number
	for i, v := range m {

		fmt.Println(v, i)
	}

	// unicode code point rune

	for i, c := range "bekar" {

		fmt.Println(i, c)

	}

	// if we want to print actual 'string'

	for i, c := range "bekar" {

		fmt.Println(i, string(c))

	}

}
