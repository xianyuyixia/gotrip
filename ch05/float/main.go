package main

import "fmt"

func main() {
	a := float32(10.0)
	b := 10.0
	c := .68
	d := 68.
	e := 10.e2
	f := 0x1.ap+0
	g := 1.625

	fmt.Printf("a的值为%f\n", a)
	fmt.Printf("b的值为%f\n", b)
	fmt.Printf("c的值为%f\n", c)
	fmt.Printf("d的值为%f\n", d)
	fmt.Printf("e的值为%f\n", e)
	fmt.Printf("f的值为%f\n", f)
	fmt.Printf("f的值为%.2f\n", f)
	fmt.Printf(".68的十进制科学计数法形式为%e\n", c)
	fmt.Printf("1.625的十六进制科学计数法形式为%x\n", g)
}
