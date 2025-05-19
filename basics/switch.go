package main

import "fmt"

func switch_go() {
	checkType(10)
	checkType("asdasd")

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("String")
	case byte:
		fmt.Println("Byte")
	default:
		fmt.Println("Not yet known type")
	}
}
