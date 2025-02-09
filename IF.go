package main

import (
	"fmt"
)

func main() {

	//if의 초기문에서 선언한 변수는 if문 밖에서는 사용할 수 없음
	if num := 3; num == 1 {
		fmt.Println(num)
	} else if num == 2 {
		fmt.Println(num)
	} else if num == 3 {
		fmt.Println(num)
	} else {
		fmt.Println("num is not 1, 2, 3")
	}
}
