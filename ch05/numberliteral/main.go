package main

import (
	"fmt"
)

func main() {
	d1 := 10      // 十进制
	d2 := 1_0     // 十进制
	d3 := 0010    // 八进制，表示十进制的8
	d4 := 00_10   // 八进制，表示十进制的8
	d5 := 0x10    // 十六进制，表示十进制的16
	d6 := 0x_10   // 十六进制，表示十进制的16
	d7 := 0b0010  // 二进制，表示十进制的2
	d8 := 0b_0010 // 二进制，表示十进制的2

	fmt.Println("d1 10的值为", d1)
	fmt.Println("d2 1_0的十进制值为", d2)
	fmt.Println("d3 0010的十进制值为", d3)
	fmt.Println("d4 00_10的十进制值为", d4)
	fmt.Println("d5 0x10的十进制值为", d5)
	fmt.Println("d6 0x_10的十进制值为", d6)
	fmt.Println("d7 0b0010的十进制值为", d7)
	fmt.Println("d8 0b_0010的十进制值为", d8)
}
