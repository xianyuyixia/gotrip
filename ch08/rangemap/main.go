package main

import "fmt"

func main() {
	usernameMap := map[string]string{
		"1": "闲渔一下",
		"2": "小明",
	}

	for key, value := range usernameMap {
		fmt.Printf("key:%s,value:%s\n", key, value)
	}

	// 如果只需要键，可以只获取一个值
	for key := range usernameMap {
		fmt.Printf("key:%s\n", key)
	}

	// 如果只需要键，可以使用空白标识符忽略掉第二个值
	for key, _ := range usernameMap {
		fmt.Printf("key:%s\n", key)
	}

	// 同样地，如果只需要值，可以使用空白标识符忽略掉第一个值
	for _, value := range usernameMap {
		fmt.Printf("value:%s\n", value)
	}
}
