package main

import (
	"fmt"
)

func main() {
	var arr [3][3]int
	arr = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(arr[0][0], arr[2][1])
}
