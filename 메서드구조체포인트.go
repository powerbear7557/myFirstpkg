package main

import (
	"fmt"
)

// 구조체 정의
type Calculator struct {
	X int
}

// Calculator 구조체의 메서드 함수 정의
func (c *Calculator) add(y int) {
	c.X += y //Calculator 구조체를 포인터로 하여 월래 값에 계산 된 값을 저장하여 값을 변경 할 수 있도록 함.
}

//만약 포인터를 사용하지 않으면 메서드 함수 내 calc c.X은 main함수의 calc c.X와 메모리 주소가 다름. 그러므로 main의 calc c.X깂은 그대로 10이고 전혀 다른
//곳에 15가 저장 됨. 그러므로 main의

func main() {
	calc := Calculator{10} //calc 구조체 정의. X에 10 들어감
	calc.add(5)            //Calculator를 포인터로 정의 했기 때문에 calc 구조체의 c.X의 값이 변경 되어 저장 됌.
	fmt.Println(calc.X)
}

//메서드 함수에서 Calculator 구조체를 포인터로 정의 사지 않았다면 main에서 calc Calculator구조체의 메모리 주소와 함수내 c Calculator구조체는
//전혀 다른 메모리 주소를 가진다. 그러므로 calc.add(5)로 메인에서 calc메서드 함수를 사용하면 calc := Calculator{10}인하여 calc.X 10은 그대로
//있다 하지만 포인터로 정의 하면 메인 calc Calculator구조체의 메모리 주소와 함수내 c Calculator구조체 메모리 주소가 같다. 그러므로 변경 저장
//이 가능하다.
