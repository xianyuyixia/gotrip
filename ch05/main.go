package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := int(5), uint(5)
	var c uintptr = 0x123456
	fmt.Println("int的长度为", unsafe.Sizeof(a))
	fmt.Println("uint的长度为", unsafe.Sizeof(b))
	fmt.Println("uintptr的长度为", unsafe.Sizeof(c))
}
