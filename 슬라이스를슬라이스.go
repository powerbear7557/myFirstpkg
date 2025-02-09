package main

import (
	"fmt"
)

func main() {
	var e1 []int
	e1 = []int{1, 2, 4, 5} // e1 := []int{1,2,3,4,5}

	e2 := e1[1:4]
	fmt.Println(e2)
	fmt.Printf("%T \n", e2)

	e3 := e1[:4]
	fmt.Println(e3)
	fmt.Printf("%T \n", e3)

	e4 := e1[1:3:4] //마지막은 e4 슬라이스의  cap 4보다 하나 작은 3
	fmt.Println(e4)
	fmt.Printf("%T \n", e4)
	fmt.Println(cap(e4))

}
