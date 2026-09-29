package main

import (
	"fmt"
	"time"
)

type order struct {
	id     string
	amount float32
	status string
	somoy  time.Time
}

func main() {

	myorder1 := order{

		id:     "shahadat",
		amount: 4000.50,
		status: "Paid",
		somoy:  time.Now(),
	}
	fmt.Println("1st customer", myorder1)
	fmt.Println(myorder1.status)

	myorder2 := order{

		id:     "Gazi",
		amount: 6000.50,
		status: "unPaid",
		somoy:  time.Now(),
	}
	fmt.Println("2nd customer", myorder2)
	fmt.Println(myorder1.status)
	myorder2.status = "unpaid"
}
