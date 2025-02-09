package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Vertex struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	v := Vertex{1, 2}
	data, _ := json.Marshal(v)

	fmt.Println(string(data))
}

// var v0 Vertex = Vertex{1, 2}
// fmt.Println(v0)
// fmt.Println(v0.X, v0.Y)

// var v Vertex = Vertex{X: 1}
// fmt.Println(v)

// v.Y = 2
// fmt.Println(v)
// fmt.Println("X = ", v.X)
// fmt.Println("Y = ", v.Y)
