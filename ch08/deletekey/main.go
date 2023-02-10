package main

import "fmt"

func main() {
	usernameMap := map[string]string{
		"1": "闲渔一下",
		"2": "小明",
	}

	fmt.Printf("usernameMap:%v\n", usernameMap)

	delete(usernameMap, "1")

	fmt.Printf("usernameMap:%v\n", usernameMap)
}
