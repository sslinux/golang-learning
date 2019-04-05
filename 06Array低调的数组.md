# 06Array低调的数组


Go 语言里面的数组其实很不常用，这是因为数组是定长的静态的，一旦定义好长度就无法更改，而且不同长度的数组属于不同的类型，之间不能相互转换相互赋值，用起来多有不方便之处。

切片是动态的数组，是可以扩充内容增加长度的数组。当长度不变时，它用起来就和普通数组一样。当长度不同时，它们也属于相同的类型，之间可以相互赋值。这就决定了数组的应用领域都广泛地被切片取代了。

不过也不可以小瞧数组，在切片的底层实现中，数组是切片的基石，是切片的特殊语法隐藏了内部的细节，让用户不能直接看到内部隐藏的数组。切片不过是数组的一个包装，给顽固的数组装上了灵活的翅膀，让石头也可以展翅飞翔。


## 数组变量的定义：

```go
package main
// 只申明类型，不赋初值。这时编译器会给数组默认赋上「零值」。
// 数组的零值就是所有内部元素的零值。

import "fmt"

func main() {
    var a[9]int
    fmt.Println(a)
}
```

另外三种变量定义的形式， 效果都是一样的

```go
package main

import "fmt"

func main() {
	var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
```

## 数组的访问：

```go
// 使用下标来简单操作一下数组，这个数组里存的是数字的平方值
package main

import "fmt"

func main() {
	var squares [9]int
	for i := 0; i < len(squares); i++ {
		squares[i] = (i + 1) * (i + 1)
	}
	fmt.Println(squares)
}
// 内置函数 len() 来直接获取数组的长度。
// 数组的长度是编译器确定的，当我们使用 len() 函数访问数组的长度属性时，编译器在背后偷偷把它替换成了整数值。
```

## 数组的下标越界检查：

```go
package main

import "fmt"

func main() {
    var a = [5]int{1,2,3,4,5}
    a[101] = 255          // 下标越界；
    fmt.Println(a)
}
```

Go 语言会对数组访问下标越界进行编译器检查。

有一个重要的问题是，如果下标是一个变量，Go 是如何检查下标越界的呢？
变量需要在运行时才可以确定是否越界，Go 是如何办到的呢？

```go
package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = 101
	a[b] = 255      // 下标明显是超出范围了的；
	fmt.Println(a)
}
```

但是，编译是可以通过的，只有在运行时才会报错；

```bash
sslinux@sslinux-devops:~/gitbook/golang-learning$ go run 06Array_IndexOutOfRange.go 
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
        /home/sslinux/gitbook/golang-learning/06Array_IndexOutOfRange.go:8 +0x11
exit status 2
sslinux@sslinux-devops:~/gitbook/golang-learning$ go build 06Array_IndexOutOfRange.go
```

* 总结： Go 会在编译后的代码中插入下标越界检查的逻辑，所以数组的下标访问效率是要打折扣的，比不得 C 语言的数组访问性能。


## 数组赋值：

`同样的子元素类型并且是同样长度`的数组才可以相互赋值，否则就是不同的数组类型，不能赋值。

数组的赋值本质上是一种`浅拷贝`操作，赋值的两个数组变量的`值不会共享`。

```go
package main 

import "fmt"

func main() {
    var a = [9]int{1,2,3,4,5,6,7,8,9}
    var b [9]int
    b = a             // 要求：数组长度一致，元素类型一致；
    a[0] = 12345      // 只会改变a[0],不影响b[0]
    fmt.Println(a)
    fmt.Println(b)
}
// 赋值后两个数组并没有共享内部元素。
// 如果数组的长度很大，那么拷贝操作是有一定的开销的，使用的时候一定需要注意。
```

下面我们尝试使用不同长度的数组赋值会有什么结果:

```go
package main

import "fmt"

func main() {
	var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b [10]int

	b = a
	fmt.Println(b)
}

// cannot use a (type [9]int) as type [10]int in assignment
```

* `不同长度的数组之间赋值是禁止的，因为它们属于不同的类型。`


## 数组的遍历

* 使用下标进行遍历
* 使用range进行遍历

range 遍历提供了下面两种形式：

```go
package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}

	// 使用下标遍历：
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println("*******************************")

	// 使用for range 1:
	for index := range a {
		fmt.Println(index, a[index])
	}

	fmt.Println("*******************************")

	// 使用for range 2:
	for index, value := range a {
		fmt.Println(index, value)
	}
}
```

