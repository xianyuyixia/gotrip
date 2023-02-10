package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 从第4个元素开始，总共取5个元素
	slice2 := slice1[3:8]
	fmt.Printf("slice2 %v，长度：%v，容量：%v", slice2, len(slice2), cap(slice2))
}
