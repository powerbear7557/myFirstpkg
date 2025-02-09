package main

import (
	"fmt"
)

func main() {
	var num int = 2 // num := 1

	//If you wonder what the num=1 is address in the memoory. You can use '&'
	fmt.Println(&num)

	var numPointer *int = &num // 타입 지정 없이 왈로스를 사용해 numPointer := &num (num의 메모리 할당 주소를 numPointer로 저장하여라라)
	// numPointer는 포인터 선언이 되어 있고 주소를 저장한다. 그러므로 &num을 통해 2가 저장 된 num의 메모리 주소를 numPointer에 저장한다.
	//그러므로 numPointer에 저장 된 주소를 numPointer를 포이터로 정의 했기 때문에 메모리의 물리적 주소인 것을 안다.
	//그러므로 포인터가 가르치는 메모리의 값은 읽는 '*'를 사용하면 값을 읽어 올 수 있다.
	// 그러므로 *numPointer 메모리부터 값을 읽어 올 수 있다. 그러나 *num으로는 물리적 주소는 같지만 읽어 올 수 없다. 왜냐하면 num은 포인터로
	//선언이 안 되고 일반 int로 선언이 되었기 때문이다.

	fmt.Println(numPointer)
	fmt.Println(*numPointer) //if you use '*', you can read a value of num form memory.

	*numPointer = 10
	fmt.Println(*numPointer)
}
