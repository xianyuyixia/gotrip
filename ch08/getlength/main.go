package main

import "fmt"

func main() {
	map1 := make(map[string]int, 5)

	fmt.Printf("map1:%v的键值对个数%d\n", map1, len(map1))

	map1["a"] = 1

	fmt.Printf("map1:%v的键值对个数%d\n", map1, len(map1))

	map1["b"] = 2

	fmt.Printf("map1:%v的键值对个数%d\n", map1, len(map1))
}
