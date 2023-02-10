package main

import "fmt"

func main() {
	map1 := map[string]int{}

	// 不存在键a,获取到的值为int的零值
	var1 := map1["a"]
	fmt.Printf("var1=%d\n", var1)

	// Comma ok
	var2, ok := map1["a"]
	fmt.Printf("var2=%d ok=%t\n", var2, ok)

	map1["a"] = 1

	var3, ok := map1["a"]
	fmt.Printf("var3=%d ok=%t\n", var3, ok)

	// 若只是想判断是否存在指定键，不关心值，可以使用空标识符
	_, ok = map1["a"]
	fmt.Printf("ok=%t\n", ok)
}
