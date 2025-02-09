package main

import "fmt"

func main() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i // sum = sum + i
	// 	fmt.Println(i)
	// }
	// fmt.Println(sum)

	// i := 1

	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i > 10 {
	// 		break
	// 	}
	// }

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
