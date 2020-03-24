package main

import "fmt"

type person struct {
	name string
	age  int
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
}
