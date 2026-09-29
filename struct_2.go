package main

import "fmt"

type order struct {
	id       string
	money    int
	position string
}

func newOrder(id string, money int, position string) *order {

	myOrder3 := order{

		id:       id,
		money:    money,
		position: position,
	}
	return &myOrder3

}

func (o *order) changePosition(position string) {
	o.position = position

}

func (o order) getMoney() int {

	return o.money
}

func main() {
	myOrder := order{
		id:       "1",
		money:    56500,
		position: "Junior Eng.",
	}

	myOrder2 := order{
		id:       "2",
		money:    90000,
		position: "Senior Eng.",
	}

	myOrder3 := newOrder("3", 99999, "MD")
	myOrder.changePosition("CEO")
	myOrder.money = 65000
	fmt.Println(myOrder)
	fmt.Println(myOrder2)
	fmt.Println(myOrder.getMoney())
	fmt.Println(myOrder3)

	fmt.Println(myOrder3.id)
	fmt.Println(*myOrder3)

	//create struct in main function

	teacher := struct {
		position string
		salary   int
	}{"Ass.Teacher", 38000}

	fmt.Printf("Employ position: %v and Salary:%v \n ", teacher.position, teacher.salary)

}
