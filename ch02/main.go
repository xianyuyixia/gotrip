// 所属包名
package main

// 引入fmt包
import "fmt"

// main函数
func main() {
	var content string = "hello world"
	var content2 = "不使用这个变量会报错"
	// 调用fmt包里的Println函数，打印content的内容到终端
	fmt.Println(content)
	// 使用空白标识符_忽略掉content2
	// 注释下面语句，会报错
	_ = content2
}
