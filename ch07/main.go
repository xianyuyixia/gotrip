package main

import "fmt"

func main() {
	// 声明不初始化，默认为数据类型的零值
	var arr1 [5]int
	fmt.Println("arr1", arr1)

	// 声明并初始化
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2", arr2)

	// 为某些位置设置值，其余保持为0
	var arr3 [5]int = [5]int{3: 9}
	fmt.Println("arr4", arr3)

	// 使用...自动设置元素个数
	var arr4 = [...]int{4: 5}
	fmt.Println("arr4", arr4)

	// 多为数组
	var arr5 = [2][3]int{[3]int{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr5", arr5)
}
