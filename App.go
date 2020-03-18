package main

import "fmt"

func main() {
	//声明变量
	fmt.Print("Hello Word!")
	//动态推断类型
	var a int = 10
	var b = 10
	//可以省略var的写法
	c := 10
	fmt.Print(a + b + c)
	g, h := 10, "lxyer"
	fmt.Print(add(2, 4))
	fmt.Print(swap(string(g), h))

}
func add(num1, num2 int) int {
	return num1 + num2
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}
