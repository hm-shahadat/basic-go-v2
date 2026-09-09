package main

import (
	"fmt"
	"maps"
)

func main() {
	//create map
	m := make(map[string]string)

	m["name"] = "shahadat"
	m["post"] = "student"

	// fmt.Println("name:", m["name"], "And position:", m["post"])

	n := make(map[string]int)
	n["age"] = 23
	n["id"] = 48095

	o := make(map[string]int)
	o["serial"] = 67
	o["year"] = 4

	// fmt.Println("age:", n["age"], "And id:", n["id"])

	// fmt.Println(len(n))

	delete(o, "serial")

	// fmt.Println(o["serial"])

	clear(o)
	// fmt.Println(o)

	f := map[string]int{"result": 96, "Place": 4}

	j, ok := f["Place"]
	fmt.Println(j)

	if ok {
		fmt.Println("it's ok")

	} else {
		fmt.Println("not ok")
	}

	//check equal or not

	a := map[string]int{"score": 2, "position": 780}
	b := map[string]int{"score": 2, "position": 780}

	fmt.Println(maps.Equal(a, b))

}
