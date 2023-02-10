package main

import "fmt"

func main() {
	var slice1 []int

	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)

	// 同时添加多个元素
	slice1 = append(slice1, 4, 5)

	fmt.Printf("slice1 %v len=%d cap=%d\n", slice1, len(slice1), cap(slice1))
}
