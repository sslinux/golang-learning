# 08Map字典

字典在数学上的词汇是映射，将一个集合中的所有元素关联到另一个集合中的部分或全部元素，并且只能是一一映射或者多对一映射。

![map](images/map.jpg)

数组切片让我们具备了可以操作一块连续内存的能力，它是对同质元素的统一管理。而字典则赋予了不连续不同类的内存变量的关联性，它表达的是一种因果关系，字典的 key 是因，字典的 value 是果。

如果说数组和切片赋予了我们步行的能力，那么字典则让我们具备了跳跃的能力。

指针、数组切片和字典都是容器型变量，字典比数组切片在使用上要简单很多，但是内部结构却无比复杂。


## 字典的创建：

在创建字典时，必须要给 key 和 value 指定类型。创建字典也可以使用 make 函数。

```go
package main

import "fmt"

func main() {
	var m map[int]string = make(map[int]string)
	fmt.Println(m, len(m))
}
```

使用 make 函数创建的字典是空的，长度为零，内部没有任何元素。如果需要给字典提供初始化的元素，就需要使用另一种创建字典的方式。

```go
package main

import "fmt"

func main() {
	var m map[int]string = map[int]string{
		90: "优秀",
		80: "良好",
		60: "及格", // 注意：这里逗号不可缺少，否则会报语法错误；
	}
	fmt.Println(m, len(m))
}
```

字典变量同样支持类型推导，上面的变量定义可以简写成:

```go
var m = map[int]string{
    90: "优秀",
    80: "良好",
    60: "及格",
}
```

如果你可以预知字典内部键值对的数量，那么还可以给 make 函数传递一个整数值，通知运行时提前分配好相应的内存。这样可以避免字典在长大的过程中要经历的多次扩容操作。

```go
var m = make(map[int]string, 16)
```

## 字典的读写：

同 Python 语言一样，字典可以使用中括号来读写内部元素，使用 delete 函数来删除元素。

```go
package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	//读取元素：
	var score = fruits["banana"]
	fmt.Println(score)

	// 增加或修改元素：
	fruits["pear"] = 3
	fmt.Println(fruits)

	// 删除元素：
	delete(fruits, "pear")
	fmt.Println(fruits)
}
```

## 字典key不存在会怎样？

删除操作时，如果对应的 key 不存在，delete 函数会静默处理。遗憾的是 delete 函数没有返回值，你无法直接得到 delete 操作是否真的删除了某个元素。你需要通过长度信息或者提前尝试读取 key 对应的 value 来得知。

读操作时，如果 key 不存在，也不会抛出异常。它会返回 value 类型对应的零值。如果是字符串，对应的零值是空串，如果是整数，对应的零值是 0，如果是布尔型，对应的零值是 false。

你不能通过返回的结果是否是零值来判断对应的 key 是否存在，因为 key 对应的 value 值可能恰好就是零值，比如下面的字典你就不能判断 "durin" 是否存在

```go
var m = map[string]int {
    "durin": 0   
}
```

这时候必须使用字典的特殊语法，如下：

```go
package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	var score, ok = fruits["durin"] //此时字典中不存在该key
	if ok {
		fmt.Println(score)
	} else {
		fmt.Println("durin not exists")
	}

	fruits["durin"] = 0
	score, ok = fruits["durin"] // 此时字典中有这个key，但是值为0；
	if ok {
		fmt.Println(score)
	} else {
		fmt.Println("durin not exists")
	}
}
```

字典的下标读取可以返回两个值，使用第二个返回值都表示对应的 key 是否存在。

正常的函数调用可以返回多个值，但是并不具备这种“随机应变”的特殊能力 —— 「多态返回值」。


## 字典的遍历：

字典的遍历提供了下面两种方式，一种是需要携带 value，另一种是只需要 key，需要使用到 Go 语言的 range 关键字。

```go
package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	for key, val := range fruits {
		fmt.Println(key, val)
	}

	for key := range fruits {
		fmt.Println(key)
	}
}
```

Go 语言的字典没有提供诸于 keys() 和 values() 这样的方法，意味着如果你要获取 key 列表，就得自己循环一下，如下:

```go
package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	var names = make([]string, 0, len(fruits))
	var scores = make([]int, 0, len(fruits))
	// var names = make([]string, len(fruits))  // 已经是满容切片，需要扩展；
	// var scores = make([]int, len(fruits))

	for name, score := range fruits {
		names = append(names, name)
		scores = append(scores, score)
	}

	fmt.Println(names, scores)
}

```

## 线程(协程)安全

Go 语言的内置字典不是线程安全的，如果需要线程安全，必须使用锁来控制。

## 字典变量里存的是什么？

字典变量里存的只是一个地址指针，这个指针指向字典的头部对象。所以字典变量占用的空间是一个字节，也就是一个指针的大小，64 位机器是 8 字节，32 位机器是 4 字节。

![map_header](images/map_header.jpg)


* `可以使用 unsafe 包提供的 Sizeof 函数来计算一个变量的大小.`

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m = map[string]int{
		"apple":  2,
		"pear":   3,
		"banana": 5,
	}
	fmt.Println(unsafe.Sizeof(m))  // 输出为 8 ， 因为死64位系统；
}
```




