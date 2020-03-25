package main

import "fmt"

type person struct {
	name string
	age  int
}
type Point struct {
	x, y int
}

//结构体
func main() {
	p := person{"test", 33}
	p.age = 22
	fmt.Println(p)
	fmt.Println(p.age)
	//数组
	var strList [2]string
	strList[0] = "hello"
	strList[1] = "world"
	intList := [3]int{1, 2, 3}
	fmt.Println(intList)
	var s []int = intList[0:2]
	fmt.Println(s)
	var m = map[string]Point{
		"a": {2, 3},
		"b": {4, 5},
	}
	delete(m, "a")

	fmt.Println(m)
}
