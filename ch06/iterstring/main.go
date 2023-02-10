package main

import "fmt"

func main() {
	s1 := "闲渔一下"

	fmt.Println("方式一")
	//方式一，遍历出来的是utf8编码的字节
	for i := 0; i < len(s1); i++ {
		fmt.Printf("0x%x ", s1[i])
	}

	fmt.Println()

	// 以字符格式输出为乱码
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i])
	}

	fmt.Println("\n方式二")

	// 方式二，遍历出来的是Unicode字符的码点值
	for i, ch := range s1 {
		fmt.Printf("偏移量%d对应的字符为%c\n", i, ch)
	}

	fmt.Println()

	// 打印每个字符的码点值,十六进制表示
	for _, ch := range s1 {
		fmt.Printf("%c:%x ", ch, ch)
	}

	fmt.Println()

	// 打印每个字符的码点值,十进制表示
	for _, ch := range s1 {
		fmt.Printf("%c:%d ", ch, ch)
	}
}
