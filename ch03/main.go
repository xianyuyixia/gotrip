package main

// 全局变量声明
var globalA1 int
var globalB1 = 1

var (
	globalC1 = 1
	globalC2 = 1
)

func main() {
	// 指定变量类型，没有初始化
	var a1 int

	// 声明并初始化
	var b1 int = 1

	// 变量声明块,批量声明多个变量
	var (
		c1 int    = 128
		c2 int8   = 6
		c3 string = "hello"
		c4 rune   = 'A'
		c5 bool   = true
	)

	// 在一行同时声明多个变量，类型一致
	var d1, d2, d3 int = 1, 2, 3

	// 在一行同时声明多个变量，类型不一致
	var (
		e1, e2, e3 = "e1", 2, 3.0
	)

	// 省略类型信息的声明,类型为默认类型
	var f1 = 1          // int
	var f1_1 = int32(1) // int32
	var f2 = 1.1        // float64
	var f3 = 'c'        // int32
	var f4 = "d"        // string
	var f5 = true       // bool

	// 短变量声明，申明并赋值，只能在函数体中使用
	// 相当于
	// var a int
	// a = 12
	g1 := 12
	g2 := 'A'
	g3 := "hello"
	g4, g5, g6 := 12, 'A', "hello"

	// 使用空白标识符_忽略掉未使用的变量
	_ = globalA1
	_ = globalB1
	_, _ = globalC1, globalC2
	_ = a1
	_ = b1
	_, _, _, _, _ = c1, c2, c3, c4, c5
	_, _, _ = d1, d2, d3
	_, _, _ = e1, e2, e3
	_, _, _, _, _, _ = f1, f1_1, f2, f3, f4, f5
	_, _, _, _, _, _ = g1, g2, g3, g4, g5, g6

	// 当有新变量申明时，需要使用短变量声明方式
	_, h7 := 0, 0
	_ = h7
}
