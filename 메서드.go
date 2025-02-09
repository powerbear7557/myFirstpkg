package main

import (
	"fmt"
)

type Calculator struct {
	X int
}

// c는 구조체가 된 건 이다. 그러므로 c.X면 Calculator struct의 X가 된다.
func (c Calculator) add(y int) int {
	return c.X + y
}

func main() {
	calc := Calculator{X: 10} //Calculator 구조체 X에 10을 넣음.
	//calc := Calculator{10}으로 해도 된다. //calc은 X값에 10이 들어간 구조체로 정의가 된 것임
	val := calc.add(20) //calc이 X값이 10을 가지는 Calculator구조체로 정의 된 것임. 그러므로 Calculator 메서드 함수인 add를 사용할 수 있음
	fmt.Println(val)

}
