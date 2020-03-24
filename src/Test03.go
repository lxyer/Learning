package main

import "fmt"

//对指针进行操作
func main() {
	i, j := 42, 2701
	p := &i
	//打印i的指针
	fmt.Println(p) //0xc0000160c0
	*p = 21
	fmt.Println(i) //21
	p = &j
	fmt.Println(p) //0xc0000160c8
	*p = *p / 37
	fmt.Println(j) //73

}
