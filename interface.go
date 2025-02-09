package main

import (
	"fmt"
)

type Ducky interface {
	DuckSound() string
	DuckWalk() string
	isSwim() string
}

type Bird struct {
	name string
}

func (b Bird) DuckSound() string {
	return "꽥꽥꽥"
} //(b Bird)는 Bird 구조체의 메서드를 정의할 때 '이 메서드는 Bird에 속해 있다.'하고 알려주는 역할.
//즉 **(b Bird)**는 메서드의 리시버(Receiver) 입니다.
//즉, Bird 구조체와 연결된 메서드를 정의하는 역할을 합니다.

func (b Bird) DuckWalk() string {
	return "뒤뚱뒤뚱"
}

func (b Bird) isSwim() string {
	if b.name == "오리" {
		return "있다"
	} else {
		return "없다"
	}
}

func main() {
	duck := Bird{name: "오리"}
	fmt.Printf("%s 는 %s 하고 울며, %s 걸으며, 수영을 할 수 %s \n", duck.name, duck.DuckSound(), duck.DuckWalk(), duck.isSwim())

}
