package main

import "fmt"

func func_go() {
	greet := func() {
		fmt.Println("Greeting!")
	}

	greet()

	operation := add

	fmt.Println(operation(3, 5))

	rs, rm := divide(9, 5)
	fmt.Println("9/5: ", rs, " ", rm)
}
func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	var result, remain int
	result = a / b
	remain = a % b
	return result, remain
}
