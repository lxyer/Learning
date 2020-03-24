package main

import (
	"fmt"
	"time"
)

//for循环
func main() {
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	defer println("推迟调用1")
	defer println("推迟调用2")
	println(getSum(4))
	t := time.Now()
	fmt.Println(t.Hour())
	for i := 0; i < 50; i++ {
		go println(i)
	}
}

func getSum(i int) int {
	//for 循环
	sum := 0
	for j := 0; j < i; j++ {
		sum += j
	}
	return sum
}
