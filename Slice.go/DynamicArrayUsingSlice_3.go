package main

import "fmt"

func main() {
	var studentsname []string
	var studentCount int
	var studentName string

	fmt.Print("Enter how many students:")
	fmt.Scan(&studentCount)

	for i := 0; i < studentCount; i++ {

		fmt.Print("Enter students name:")
		fmt.Scan(&studentName)
		studentsname = append(studentsname, studentName)
	}
	fmt.Println(studentsname)
}
