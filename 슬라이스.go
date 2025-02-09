package main

import (
	"fmt"
)

func main() {
	var s []int  //슬라이스 선언
	var s1 []int //슬라이스 선언

	//슬라이스 초기화 2가지 방법
	s = []int{1, 2, 3}
	s1 = make([]int, 3)
	s2 := []int{6, 7, 8, 9, 10}
	s3 := make([]int, 10)

	//슬라이스 읽기
	fmt.Println(s[0])
	fmt.Println(s1[0])
	fmt.Println(s1[1])
	fmt.Println(s1[2])

	//슬라이스 값 저장
	s[0] = 10
	s1[0] = 100

	fmt.Println(s[0])
	fmt.Println(s1[0])
	fmt.Println(s2[0], s2[1], s2[3])
	fmt.Println(s3)

}
