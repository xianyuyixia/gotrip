package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	fmt.Printf("slice1 %v, cap=%d, len=%d\n", slice, cap(slice), len(slice))

	// 所需cap=5，oldCap=3，oldCap*2=6，6>5，newCap=5，但实际上新分配出来的容量是6，这是Go内存分配的机制导致的
	slice1 := append(slice, 4, 5)

	fmt.Printf("slice1 %v, cap=%d, len=%d\n", slice1, cap(slice1), len(slice1))
}
