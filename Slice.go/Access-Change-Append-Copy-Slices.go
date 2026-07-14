package main

import "fmt"

func main() {

	//Access Elements of a Slice
	s1 := []int{10, 20, 30}
	fmt.Println("Access Elements of a Slice", s1[2])

	// Change Elements of a Slice
	s1[2] = 3
	fmt.Println("Change Elements of a Slice", s1[2])

	s1[2] = 30
	fmt.Println("Change Elements of a Slice", s1[2])

	// Append Elements To a Slice
	s1 = append(s1, 40, 50)
	fmt.Println("Append Elements To a Slice", s1)
	fmt.Printf("length: %d\n", len(s1))
	fmt.Printf("capacity: %d\n", cap(s1))

	// Append One Slice To Another Slice
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6}
	s4 := append(s2, s3...)

	fmt.Printf("slice s4: %v\n", s4)
	fmt.Printf("length s4: %d\n", len(s4))
	fmt.Printf("Capacity s4: %d\n", cap(s4))

	// Change The Length of a Slice
	ar1 := []int{0, 00, 000, 0000, 00000}
	sl1 := ar1[1:4]
	fmt.Printf("slice s4: %v\n", sl1)
	fmt.Printf("length s4: %d\n", len(sl1))
	fmt.Printf("Capacity s4: %d\n", cap(sl1))

	sl1 = ar1[1:3]
	fmt.Printf("slice s4: %v\n", sl1)
	fmt.Printf("length s4: %d\n", len(sl1))
	fmt.Printf("Capacity s4: %d\n", cap(sl1))

	sl1 = append(sl1, 1, 2, 3)
	fmt.Printf("slice s4: %v\n", sl1)
	fmt.Printf("length s4: %d\n", len(sl1))
	fmt.Printf("Capacity s4: %d\n", cap(sl1))

	sl1 = append(sl1, 5, 6, 7)
	fmt.Printf("slice s4: %v\n", sl1)
	fmt.Printf("length s4: %d\n", len(sl1))
	fmt.Printf("Capacity s4: %d\n", cap(sl1))

	//  The copy() function
	sl2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice s4: %v\n", sl2)
	fmt.Printf("length s4: %d\n", len(sl2))
	fmt.Prin~tf("Capacity s4: %d\n", cap(sl2))

	sl3 := sl2[:len(sl2)-4]      //6
	sl4 := make([]int, len(sl3)) //6
	copy(sl4, sl3)

	fmt.Printf("value: %v\n", sl4)
	fmt.Printf("length: %d\n", len(sl4))
	fmt.Printf("capacity: %d\n", cap(sl4))
}
