# Go语言学习笔记-08map数据类型

Go语言中的`map`，类似于其它语言的映射，字典，哈希表，表示一组无序的键值对

## 声明

```go
var mapVar map[key_type]value_type
```

`key_type`为键的类型，不能是函数类型，map类型，切片类型

`value_type`为值的类型

没有显示初始化的`map`类型，默认值为`nil`，是无法使用的，可以通过以下两种方式初始化

声明同时初始化

```go
var map1 = map[key_type]value_type{}
```

使用`make`函数初始化，可以指定容量，但不会受限于初始的容量，`map`会在键值对的个数超过当前长度时，会动态扩容

```go
var map1 = make(map[key_type]value_type, 容量)
```

ch08/main.go

```go
package main

import "fmt"

func main() {
	var map1 map[string]int

	// 无法赋值，map1尚未初始化
	//map1["a"] = 1
	fmt.Printf("map1 %#v\n", map1)

	map1 = map[string]int{}

	map1["a"] = 1

	fmt.Printf("map1 %v\n", map1)

	map2 := make(map[string]int, 5)

	fmt.Printf("map2 %v", map2)
}
```

输出

```
map1 map[string]int(nil)
map1 map[a:1]
map2 map[]  
```

## 获取键值对个数

使用内置函数`len`获取`map`的键值对个数

ch08/getlength/main.go

```go
package main

import "fmt"

func main() {
	map1 := make(map[string]int, 5)

	fmt.Printf("map1:%v的键值对个数%d\n", map1, len(map1))

	map1["a"] = 1

	fmt.Printf("map1:%v的键值对个数%d\n", map1, len(map1))

	map1["b"] = 2

	fmt.Printf("map1:%v的键值对个数%d\n", map1, len(map1))
}
```

输出

```
map1:map[]的键值对个数0
map1:map[a:1]的键值对个数1    
map1:map[a:1 b:2]的键值对个数2
```

## 读取数据

获取一个不存在的键数据时，返回的是该值数据类型的零值

通过`Comma ok`语法格式可以判断是否存在某个键，该形式如下：

```go
v,ok := mapVar[key]
```

其中`ok`是布尔值，存在该`key`返回`true`，否则返回`false`

ch08/getvalue/main.go

```go
package main

import "fmt"

func main() {
	map1 := map[string]int{}

	// 不存在键a,获取到的值为int的零值
	var1 := map1["a"]
	fmt.Printf("var1=%d\n", var1)

	// Comma ok
	var2, ok := map1["a"]
	fmt.Printf("var2=%d ok=%t\n", var2, ok)

	map1["a"] = 1

	var3, ok := map1["a"]
	fmt.Printf("var3=%d ok=%t\n", var3, ok)

	// 若只是想判断是否存在指定键，不关心值，可以使用空标识符
	_, ok = map1["a"]
	fmt.Printf("ok=%t\n", ok)
}
```

输出

```
var1=0
var2=0 ok=false
var3=1 ok=true
ok=true 
```

## 删除数据

使用内置函数`delete`删除某个键，即使该键不存在

ch08/deletekey/main.go

```go
package main

import "fmt"

func main() {
	usernameMap := map[string]string{
		"1": "闲渔一下",
		"2": "小明",
	}

	fmt.Printf("usernameMap:%v\n", usernameMap)

	delete(usernameMap, "1")

	fmt.Printf("usernameMap:%v\n", usernameMap)
}
```

输出

```
usernameMap:map[1:闲渔一下 2:小明]
usernameMap:map[2:小明]
```

## 遍历

ch08/rangemap/main.go

```go
package main

import "fmt"

func main() {
	usernameMap := map[string]string{
		"1": "闲渔一下",
		"2": "小明",
	}

	for key, value := range usernameMap {
		fmt.Printf("key:%s,value:%s\n", key, value)
	}

	// 如果只需要键，可以只获取一个值
	for key := range usernameMap {
		fmt.Printf("key:%s\n", key)
	}

	// 如果只需要键，可以使用空白标识符忽略掉第二个值
	for key, _ := range usernameMap {
		fmt.Printf("key:%s\n", key)
	}

	// 同样地，如果只需要值，可以使用空白标识符忽略掉第一个值
	for _, value := range usernameMap {
		fmt.Printf("value:%s\n", value)
	}
}
```

输出

```
key:1,value:闲渔一下
key:2,value:小明
key:1           
key:2           
key:1           
key:2           
value:闲渔一下  
value:小明   
```

需要注意的是，我们不能依赖`map`的遍历顺序，因为`map`里面的键值对是无序的

## 笔记地址

github：https://github.com/xianyuyixia/gotrip

## 交流学习

微信搜索『闲渔一下』或『xianyyxia』关注公众号

个人wx『xianyuyixia』