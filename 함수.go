package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func calc(x int, y int) (int, int) {
	return x + y, x - y
}

func myName() string {
	name := "성욱"
	return name //함수 안에서 선언 된 이 name은 main에서 사용 할 수 없음
}

func main() {
	num := add(42, 13)
	fmt.Println(num)

	fmt.Println(add(23, 56))

	num1, num2 := calc(42, 13)
	fmt.Println(num1, num2)
	fmt.Println(calc(42, 13))
}
