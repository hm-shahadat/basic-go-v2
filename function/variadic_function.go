package main

import "fmt"

func variD(nums ...int) int {

	total := 0

	for _, num := range nums {

		total = total + num

	}
	return total

}

func main() {
	result := variD(3, 4, 5, 6)
	fmt.Println(result)
}
