package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr = [5]int{1, 2, 3, 4, 5} // arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println(slice)
	fmt.Printf("%T \n", arr)
	fmt.Printf("%T \n", slice)
}
