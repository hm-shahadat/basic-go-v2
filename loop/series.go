package main

import (
	"fmt"
)

//series program sum like 1+2+3...+10= ?

func main() {
	var startnum, endnum int
	fmt.Printf("please enter a starting number:")
	fmt.Scan(&startnum)

	fmt.Printf("please enter a ending number:")
	fmt.Scan(&endnum)
	
	sum := 0
	for i := startnum; i <= endnum; i++ {
		sum = sum + i

	}

	fmt.Println("The sum of the series:", sum)

}
