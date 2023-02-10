package main

import "fmt"

// 每新增一行，iota自增1
const (
	a1 = iota //0
	a2 = iota //1
)

// 每一个const代码块，iota重置为0
const (
	b1 = iota //0
)

const (
	failed  = iota //0
	success        //省略赋值，与上一个赋值一致，则为iota，同时新增了一行，则值为1
)

const (
	zero  = iota // 0
	one          // 1
	five  = 5    // 5
	three = iota // 行号为3，则iota的值为3
)

func main() {
	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)
	fmt.Println("b1 = ", b1)
	fmt.Println("failed = ", failed)
	fmt.Println("success = ", success)

	fmt.Println("zero = ", zero)
	fmt.Println("one = ", one)
	fmt.Println("five = ", five)
	fmt.Println("three = ", three)
}
