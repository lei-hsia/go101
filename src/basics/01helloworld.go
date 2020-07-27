package main

import "fmt"

func main() {
	fmt.Println("Go: hello world")

	// var: 声明，然后用; 文件变量声明之后可以不用，全局变量声明之后必须使用
	name = "someone"
	age = 18
	isOK = true
	//fmt.Println(isOK)
	fmt.Printf("name: %s", name)
	fmt.Print(age)

	a := 1 // short var declaration
	fmt.Print(a)

	_ = 1 // anomynous var
}
