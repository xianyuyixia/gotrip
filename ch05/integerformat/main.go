package main

import "fmt"

func main() {
	a := 255

	fmt.Printf("255的二进制表示：%b\n", a)  // 11111111
	fmt.Printf("255的十进制表示：%d\n", a)  // 255
	fmt.Printf("255的八进制表示：%o\n", a)  // 377
	fmt.Printf("255的八进制表示：%O\n", a)  // 0o377
	fmt.Printf("255的十六进制表示：%x\n", a) // ff
	fmt.Printf("255的十六进制表示：%X\n", a) // FF
}
