package main

import "fmt"

//同时申明多个变量
var i, j = 1, 2
var good, boy, age = true, false, 18

func main() {
	a, b := Swap("a", "b")
	fmt.Print(split(6))
	println(a, b)

}

//函数
func Swap(x, y string) (string, string) {
	return y, x
}

//直接return,会返回函数定义好的返回值类型
func split(sum int) (x, y int) {
	x = sum / 3
	y = sum - x
	return
}
