package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// utf-8编码，一个中文由3个字节组成，所以长度为12
	s1 := "闲渔一下"

	fmt.Printf("s1的长度为:%d\n", len(s1))
	fmt.Printf("s1的字符个数为:%d\n", utf8.RuneCountInString(s1))
}
