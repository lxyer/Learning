package main

import (
	"fmt"
)

func main() {
	//声明变量
	fmt.Print("Hello Word!")
	//动态推断类型
	var a int = 10
	var b = 10
	//可以省略var的写法
	c := 10
	fmt.Println(a + b + c)
	g, h := 10, "lxyer"
	fmt.Println(add(2, 4))
	fmt.Println(swap(string(g), h))
	var ip *int
	ip = &c
	fmt.Print(ip)
	//使用结构体
	var p1, p2 Person
	p1.name = "Steven"
	p1.age = 22
	p2.name = "Job"
	p2.age = 21
	fmt.Print(p1.name)
	//创建map
	var testMap map[string]int
	testMap["age"] = 22
	testMap["city"] = 010
	//遍历map
	for key := range testMap {
		fmt.Print(testMap[key])
	}
}

//定义结构体
type Person struct {
	name string
	age  int
}

func add(num1, num2 int) int {
	return num1 + num2
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}
