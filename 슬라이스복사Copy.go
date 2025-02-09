package main

import (
	"fmt"
)

func main() {
	var s1 []int
	s1 = []int{1, 2, 3, 4, 5} // s1 := []int{1, 2, 3, 4, 5}

	var s2 []int
	s2 = make([]int, 10) // s2 := make([]int, 10)

	copy(s2, s1) //s2가 저장 되는 변수, s1이 저장할 값
	fmt.Println(s2)

	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3 //s4 슬라이스는 s3를 참조 해서 이다
	s4[0] = 99
	fmt.Println("s3 =", s3)
	fmt.Println("s4 =", s4)

	s5 := []int{1, 2, 3, 4, 5}
	s6 := make([]int, 5)
	fmt.Println("카피전 s6 값 = ", s6)

	copy(s6, s5) // copy 명령어를 사용하면 복사를 해도 참조가 안됌. 그러므로 복사한 원본 값은 변화하지 않음
	fmt.Println("카피전 s6 값 = ", s6)

	s6[0] = 100
	fmt.Println(s6)

}
