package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Printf("slice 的长度为：%d,容量为：%d", len(slice), cap(slice))
}
