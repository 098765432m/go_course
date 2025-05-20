package main

import "fmt"

func main() {
	age := 16
	name := "Phu"
	s := fmt.Sprint("Hello", " World ", age, " ", name)

	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World", age, name)
	fmt.Print(s)
	fmt.Print(s)
}
