package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}

	slice := arr[1:5]

	fmt.Printf("slice %v, cap=%d len=%d\n", slice, cap(slice), len(slice))

	// 修改slice
	slice[0] = 22
	fmt.Printf("slice %v, cap=%d len=%d\n", slice, cap(slice), len(slice))
	// 原数组一起变化
	fmt.Printf("arr %v\n", arr)

	// 使slice扩容
	slice = append(slice, 6, 7)
	// 扩容后再修改
	slice[0] = 2
	fmt.Printf("slice %v, cap=%d len=%d\n", slice, cap(slice), len(slice))
	// 原数组无变化
	fmt.Printf("arr %v\n", arr)
}
