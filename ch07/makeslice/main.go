package main

import "fmt"

func main() {
	slice := make([]int, 5)
	slice2 := make([]int, 5, 10)

	fmt.Printf("slice %v 长度：%d，容量：%d\n", slice, len(slice), cap(slice))
	fmt.Printf("slice2 %v 长度：%d，容量：%d\n", slice2, len(slice2), cap(slice2))
}
