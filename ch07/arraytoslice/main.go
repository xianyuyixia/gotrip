package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arr[2:7:8] // [2,3,4,5,6]
	slice2 := arr[:]     // [1,2,3,4,5,6,7,8,9,10]
	slice3 := arr[3:4]   // [4]

	fmt.Printf("slice1 %v，长度：%d，容量：%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 %v，长度：%d，容量：%d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3 %v，长度：%d，容量：%d\n", slice3, len(slice3), cap(slice3))

	// 修改切片的第4个元素，及把6改为12
	slice1[3] = 12

	// 原数组arr发生改变
	fmt.Printf("修改后slice1 %v，长度：%d，容量：%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("修改后slice2 %v，长度：%d，容量：%d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("修改后arr %v\n", arr)

}
