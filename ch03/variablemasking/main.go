package main

import "fmt"

var a = 1

func main() {
	{
		// 遮蔽包级变量a
		a := 2
		fmt.Println(a) // 2
	}

	if a := 3; a == 3 {
		fmt.Println(a) // 3
	}

	// 等价于上面的if语句，外层的{}就是隐藏的代码块
	{
		a := 3
		if a == 3 {
			fmt.Println(a) // 3
		}
	}

	// 都不影响包级变量a
	fmt.Println(a) // 1
}
