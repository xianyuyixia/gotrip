package main

import "fmt"

func main() {
	s1 := "闲渔一下"

	// 使用下标获取
	fmt.Printf("s1的第一个字节为：%x\n", s1[0])
	// 取前3个字节组成的字符串，因为是utf-8编码，1个中文用3个字节表示，所以为闲
	fmt.Printf("s1的第一个字符为：%s\n", s1[0:3])
}
