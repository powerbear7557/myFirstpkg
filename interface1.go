package main

import (
	"fmt"
)

func typeCheck(i interface{}) {
	fmt.Printf("%T \n", i) //i의 타입이 무었인지 출력하는 출력문
}

func main() {
	var num1 int8 = 1
	var num2 int16 = 2
	var num3 int32 = 3
	var num4 int64 = 4
	typeCheck(num1)
	typeCheck(num2)
	typeCheck(num3)
	typeCheck(num4)
}
