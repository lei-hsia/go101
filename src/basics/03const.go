package main

import "fmt"

const (
	a1 = 100
	a2
	a3
)

// iota: const关键字出现时重置为0，const中每新增一行iota加1
const (
	n1 = iota
	n2
	_
	n3
)

const (
	b1 = iota
	b2 = 100
	b3 // 不写默认和上一行相同
	b4
	b5 = iota
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(b5)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}
