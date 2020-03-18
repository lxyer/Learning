package main

import "fmt"

func main() {
	fmt.Print("Hello Word!")
	//动态推断类型
	var a int = 10
	var b = 10
	//可以省略var的写法
	c :=10
	fmt.Print(a+b+c)

}
