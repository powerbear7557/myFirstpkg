package main

import (
	"fmt"
)

// numPointer
func addOne(num *int) {
	*num += 1 //num변수 메모리 주소에 num + 1 값을 저장하라.
	//메인함수에서 addOne(&num)으로 num 포인터의 num주소를 저장해 두었음
	//그러므로 *num의 값은 num 포인터에 저장 된 메모리 주소의 값이므로
	//*num = *num +1 ,즉 *num += 1은 num에 저장 된 메모리 주소의 값에 1을 더하고 다시 그 주소에 다시 저장하라는 의미임
	//num 주소의 초기값에 1이 저장 되었으므로 *num = 1(num 주소의 초기값) + 1 즉, num 포인트에 저장 된 메모리 주소로 2를 다시 저장한다.

	fmt.Println(&num)
	fmt.Println(num)
}

func main() {
	num := 1
	fmt.Println(&num)
	addOne(&num) //addOne function에
	fmt.Println(num)
}

//여기서 가장 중요한 것 메인과 addOne함수의 num의 메모리 물리적 주소는 다르다. 하지만 addOne()에서 num포인트에 메일의 num 주소값이
//저장 되었기 때문에 addOne()에서 메인의 num 값에 +1을 더하여 2로 저장 할 수 있었던 것이다.
