package main

import (
	"fmt"
	"math"
)

//方法
func main() {
	ver := Vertex{4, 4}
	fmt.Println(ver.Abs())
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
