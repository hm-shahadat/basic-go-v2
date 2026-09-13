package main

import "fmt"

// we can also use 'interface{}' for use any types of value
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

	// 	we can also use slice instead of 'result'
	slic := []int{3, 4, 5, 6, 7}
	result2 := variD(slic...)
	fmt.Println(result2)

}
