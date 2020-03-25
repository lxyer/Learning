package main

//闭包
func main() {
	num := addOne()
	println(num())
	println(num())
	println(num())
}

func addOne() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
