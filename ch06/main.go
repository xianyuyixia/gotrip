package main

import "fmt"

func main() {
	// 本质上是int32类型
	var ch1 rune
	ch1 = '闲'

	// 使用%c格式化输出字符
	fmt.Printf("闲的Unicode码点值：%d,字符：%c\n", ch1, ch1)

	ch2 := '渔'
	fmt.Printf("渔的Unicode码点值：%d,字符：%c\n", ch2, ch2)

	// 使用八进制数表示字符a
	ch3 := '\141'
	// 使用十六进制数表示字符a
	ch4 := '\x61'

	fmt.Printf("八进制141表示的字符为:%c\n", ch3)
	fmt.Printf("十六进制61表示的字符为:%c\n", ch4)

	// 使用\u接四个十六进制数表示
	ch5 := '\u95f2'
	fmt.Printf("\\u95f2表示的字符为：%c\n", ch5)

	// 使用\U接四个或八个十六进制数表示
	ch6 := '\U000095f2'
	fmt.Printf("\\u95f2表示的字符为：%c\n", ch6)

	// 使用byte表示单字节字符
	var ch7 byte
	ch7 = 97
	fmt.Printf("97对应的ascii为%c\n", ch7)

	// 使用rune表示一个Unicode码点
	var ch8 rune
	ch8 = 38386
	fmt.Printf("38386对应的Unicode字符为%c", ch8)
}
