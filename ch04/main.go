package main

import "fmt"

const a1 = 1

// 批量定义
const (
	b1 = true
	b2 = "a"
)

// 只定义名称，保持与上一个常量一致
const (
	c1 = 1 //1
	c2     //1
)

func main() {
	fmt.Println("a1 =", a1)

	fmt.Println("b1 =", b1)
	fmt.Println("b2 =", b2)

	fmt.Println("c1 =", c1)
	fmt.Println("c2 =", c2)
}
