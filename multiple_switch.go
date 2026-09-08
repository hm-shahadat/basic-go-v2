package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {

	case time.Friday, time.Saturday:

		fmt.Println("weekend")

	default:
		fmt.Println("workday")
	}
}
