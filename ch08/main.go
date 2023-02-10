package main

import "fmt"

func main() {
	var map1 map[string]int

	// 无法赋值，map1尚未初始化
	//map1["a"] = 1
	fmt.Printf("map1 %#v\n", map1)

	map1 = map[string]int{}

	map1["a"] = 1

	fmt.Printf("map1 %v\n", map1)

	map2 := make(map[string]int, 5)

	fmt.Printf("map2 %v", map2)
}
