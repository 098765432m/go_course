package main

import "fmt"

func format_verb() {
	// x := 255

	// fmt.Printf("%b\n", x)
	// fmt.Printf("%d\n", x)
	// fmt.Printf("%+d\n", x)
	// fmt.Printf("%o\n", x)
	// fmt.Printf("%O\n", x)
	// fmt.Printf("%x\n", x)
	// fmt.Printf("%X\n", x)
	// fmt.Printf("%#x\n", x)
	// fmt.Printf("%4d\n", x)
	// fmt.Printf("%-4d\n", x)
	// fmt.Printf("%04d\n", x)

	// str := "Hello"

	// fmt.Printf("%s\n", str)
	// fmt.Printf("%q\n", str)
	// fmt.Printf("%8s\n", str)
	// fmt.Printf("%-8s asads\n", str)
	// fmt.Printf("%x\n", str)
	// fmt.Printf("%X\n", str)

	// t := true
	// fmt.Printf("%t\n", t)

	flt := 918.0000355000

	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)
	fmt.Printf("%+v\n", flt)
}
