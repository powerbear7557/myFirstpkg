package main

import (
	"fmt"
) // import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s, 7)
	fmt.Println(s)

	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 2, 3, 4)

	fmt.Println(numbers)
}
