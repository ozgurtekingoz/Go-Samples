package main

import "fmt"

func main() {

	fmt.Println("hello world")

	var h string = "ozgur"
	fmt.Printf("h is type %T\n", h)

	fmt.Println(add(3, 5))
	fmt.Println(subt(5, 3))

	//birden fazla değer return edilebilir
	x, y := mult("hello", "world")
	fmt.Println(x, y)

	//for
	z := 1
	for z <= 5 {
		fmt.Println(z)
		z++
	}
	for z := 1; z <= 5; z++ {
		fmt.Println(z)
	}

	ifStatement()
}

func ifStatement() {
	for y := 1; y <= 5; y++ {
		if y%2 == 0 {
			fmt.Println(y, "even number")
		} else {
			fmt.Println(y, "odd number")
		}
	}
}

func add(x int, y int) int {
	return x + y
}

func subt(a, b int) int {
	return a - b

}

func mult(a, b string) (string, string) {
	return a, b
}
