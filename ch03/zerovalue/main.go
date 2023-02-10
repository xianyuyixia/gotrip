package main

import "fmt"

func main() {
	var var1 int
	var var2 float32
	var var3 bool
	var var4 string
	var var5 byte
	var var6 [5]int
	var var7 chan int
	var var8 map[string]int
	var var9 []int
	var var10 func() int

	fmt.Printf("int的零值为%v\n", var1)
	fmt.Printf("float32的零值为%v\n", var2)
	fmt.Printf("bool的零值为%v\n", var3)
	fmt.Printf("string的零值为%#v\n", var4)
	fmt.Printf("byte的零值为%v\n", var5)
	fmt.Printf("[5]int的零值为%v\n", var6)

	fmt.Printf("[]int的零值为%#v\n", var9)
	fmt.Printf("chan int的零值为%#v\n", var7)
	fmt.Printf("map的零值为%#v\n", var8)
	fmt.Printf("func() int的零值为%#v\n", var10)
}
