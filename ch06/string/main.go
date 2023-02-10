package main

import "fmt"

func main() {
	s1 := "闲渔一下"
	// 使用\u前缀接四个十六进制数
	s2 := "\u95f2\u6e14\u4e00\u4e0b"
	// 使用\U前缀接八个十六进制数
	s3 := "\U000095f2\U00006e14\U00004e00\U00004e0b"
	// 使用十六进制，utf-8编码
	s4 := "\xe9\x97\xb2\xe6\xb8\x94\xe4\xb8\x80\xe4\xb8\x8b"
	// 混合使用
	s5 := "闲\u6e14\U00004e00\xe4\xb8\x8b"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}
