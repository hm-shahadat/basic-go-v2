package main

import "fmt"

type stu1 struct {
	name    string
	id      int
	age     int
	address string
	Mamla   int
}

func (x *stu1) incriseAge(val int) {

	x.age = x.age + val

}
func (y *stu1) incriseCase(val int) {

	y.Mamla = y.Mamla + val
}
func ds(s stu1) {
	fmt.Printf("Name= %v, Id= %v,age= %v, Address= %v, CaseCount= %v\n", s.name, s.id, s.age, s.address, s.Mamla)
}

func main() {

	Sadu := stu1{"sadu", 01, 23, "Chandpur", 17}
	baba := stu1{"pirsab", 02, 24, "Chattagram", 20}

	Sadu.incriseAge(2)
	baba.incriseCase(3)
	fmt.Println("Sadu Profile: ")
	ds(Sadu)
	fmt.Println("Baba Profile: ")
	ds(baba) 

}
