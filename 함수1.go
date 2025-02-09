package main

import (
	"fmt"
)

// func addOne(num int) {
// 	num += 1
// 	fmt.Println("addOne num =", num)
// }

func main() {
	num := 1
	addOne(num)
	fmt.Println("main num =", num) //2로 과연 출력이 될까? 아닙니다. 함수와 메인 함수의 num변수는 다릅니다. 그러므로 1이 출력 됩니다.

}
