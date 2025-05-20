package main

import "fmt"

func closure() {
	// seq := adder()
	// fmt.Println(seq())
	// fmt.Println(seq())
	// fmt.Println(seq())
	// fmt.Println(seq())

	// seq2 := adder()

	// fmt.Println(seq2())

	substractor := func() func(int) int {
		countDown := 99

		return func(x int) int {
			countDown -= x
			return countDown
		}
	}()

	fmt.Println(substractor(1))
	fmt.Println(substractor(4))
	fmt.Println(substractor(84))
}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
