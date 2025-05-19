package main

import (
	"fmt"
	"slices"
)

func slice_go() {
	// var numbers1 []int
	// var numbers2 = []int{1, 2, 3}
	// numbers3 := []int{6, 3}

	// numbers4 := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	slice := a[1:4]

	slice = append(slice, 6, 7)

	fmt.Println(slice)

	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	var nilSlice = []int{}

	if slices.Equal(nilSlice, nil) {
		fmt.Println("Equal")
	}

	slice1 := slice[2:4]
	fmt.Println(slice1)
}
