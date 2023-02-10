package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 []int
	fmt.Printf("slice1 %v 长度为：%d，容量为：%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 %#v 长度为：%d，容量为：%d\n", slice2, len(slice2), cap(slice2))
}
