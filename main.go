package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Printf("%T \n", arr)
	fmt.Println(len(arr))
	fmt.Println(arr[0], arr[1], arr[2])
}
