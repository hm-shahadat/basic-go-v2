package main

import "fmt"

func main() {
	// var name string
	// const COUNTRY = "Bangladesh"
	// var age, num1, num2 int
	// var gpa float32

	// fmt.Printf("Enter your name: ")
	// fmt.Scanf("%v", &name)

	// fmt.Printf("Enter your age: ")
	// fmt.Scan(&age)

	// fmt.Printf("Enter your SSC gpa: ")
	// fmt.Scanln(&gpa)
	// fmt.Printf("Enter 2 numbers: ")
	// fmt.Scan(&num1, &num2)

	// fmt.Printf("%v is a student\n", name)
	// fmt.Printf("%v is %v years old\n", name, age)
	// fmt.Printf("%v has got GPA %v/5 in SSC\n", name, gpa)
	// fmt.Printf("%v originally from %v\n", name, COUNTRY)
	// fmt.Printf("num1 = %v, num2 = %v\n", num1, num2)
	// //

	var name string
	var age int
	var uni string
	var year int

	fmt.Println("enter a name:")
	fmt.Scanf(&name)

	fmt.Println("enter your age")
	fmt.Scanf(&age)

	fmt.Println("\"enter your uni. name\"")
	fmt.Scanf(&uni)

	fmt.Println("enter your \\passing\\ year")
	fmt.Scanf(&year)

	fmt.Printf("%v is a student and age %v\n", name, age)
	fmt.Printf("you are graduate from %v and your passing year %v\n", uni, year)

}
