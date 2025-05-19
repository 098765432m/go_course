package main

import "fmt"

func map_go() {
	myMap := make(map[string]int)

	myMap["code"] = 18
	myMap["lee"] = 28

	delete(myMap, "code")
	fmt.Println(myMap)

	myMap["key1"] = 1
	myMap["key2"] = 5
	myMap["key4"] = 528

	fmt.Println(myMap)

	value, ok := myMap["key4"]
	fmt.Println(value)
	fmt.Println(ok)

	for k, v := range myMap {
		fmt.Printf("k: %s,v = %d\n", k, v)
	}
}
