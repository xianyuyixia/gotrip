package main

import "fmt"

// 由使用的地方决定类型
const (
	one = 1 // 1
	two     // 2
)

const (
	Spring1 int64 = iota
	// Summer1 保持与Spring1的定义一致
	Summer1
)

const (
	Spring2 int64 = iota
	// Summer2 与Spring2的类型不一致，是由使用的地方决定类型
	Summer2 = iota
)

func main() {
	var a1 = one
	var a2 int32 = one
	var a3 int32 = two

	fmt.Printf("a1的类型为%T\n", a1)
	fmt.Printf("a2的类型为%T\n", a2)
	fmt.Printf("a3的类型为%T\n", a3)

	fmt.Printf("Spring1的类型为%T\n", Spring1)
	fmt.Printf("Summer1的类型为%T\n", Summer1)

	fmt.Printf("Spring2的类型为%T\n", Spring2)
	fmt.Printf("Summer2的类型为%T\n", Summer2)
}
