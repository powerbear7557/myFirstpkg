package main

import (
	"fmt"
)

func main() {
	/* 아래보다 더 간단하게 초기화하는 법
	var numMap = map[string]int{
		"one" : 1,
		"two" : 2,
		"three" : 3,
		"four" : 4,
	}
	*/
	/*
		var numMap = make(map[string]int)
		numMap["one"] = 1
		numMap["two"] = 2
		numMap["three"] = 3
		numMap["four"] = 4

		fmt.Println(numMap)
		fmt.Println(numMap["one"], numMap["two"], numMap["three"], numMap["four"], numMap["five"])
	*/

	var a = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(a["one"])
	//delete(a, "one")
	//fmt.Println(a["one"])
	//value, isKey := a["one"] //이 map에 값이 있는지? 그리고 있으면 iskey -> true
	//fmt.Println(svalue, isKey)

	delete(a, "one")
	fmt.Println(a["one"])

	value, isKey := a["one"]
	fmt.Println(value, isKey)
}
