package main

// func under pera mater
func pos(fn func(a int) int) {

	fn(1)
}

//'func' use in return type

func pos2() func(a int) int {

	return func(a int) int {
		return 4
	}
}
func main() {

	fn := func(a int) int {

		return 2
	}

	pos(fn)

	fm := pos2()
	fm(6)

}
