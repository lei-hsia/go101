package main

import "fmt"

// 函数外每个语句都需要关键字开始
var (
	name string
	age  int8
	isOK bool
)

func main() {
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOK)
}
