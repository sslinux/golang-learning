# 基本数据类型和操作符

## 定义变量的三种方式：

* 方式一：

```go
package main

import "fmt"

func main() {
    var s int = 42
    fmt.Println(s)
}
```

* 方式二： 省略变量类型；

```go
package main

import "fmt"

func main() {
    var s = 42     # 初始化时省略变量类型，自动推导类型；
    fmt.Println(s)
}
```

* 方式三： 去掉var关键字；  ==> 只能用于函数中，不能用于定义全局变量；

```go
package main

import "fmt"
func main() {
    s := 42
    fmt.Println(s)
}
```

类型是变量身份的象征，如果一个变量不那么在乎自己的身份，那在形式上就可以随意一些。

var 的意思就是告诉读者「我很重要，你要注意」， := 的意思是告诉读者「我很随意，别把我当回事」。

var 再带上显式的类型信息是为了方便读者快速识别变量的身份。

如果一个变量很重要，建议使用第一种显示声明类型的方式来定义，比如全局变量的定义就比较偏好第一种定义方式。

如果要使用一个不那么重要的局部变量，就可以使用第三种。比如循环下标变量：

```go
for i:=0; i<10; i++ {
    doSomething()
}
```

```go
var i = 0
for ; i<10; i++ {
    doSomething()
}
```


`在第一种声明变量的时候不赋初值，编译器会自动赋予相应类型的[零值]`

```go

字符串的零值是 空串，而不是nil；
整型的零值是 0 ；
布尔类型的零值是false；
```

```go
//演示定义变量时不赋初值，则变量由编译器自动赋予相应类型的零值；
package main

import "fmt"

func main() {
    var i int
    var s string
    var b bool

    fmt.Println(i)
    fmt.Println(s)
    fmt.Println(b)
}
```


## 全局变量和局部变量

局部变量：定义在函数中，函数调用结束它的就消亡了。

全局变量：定义在函数的外面，在程序运行期间，它一直存在；


```go
package main

import "fmt"

var globali int = 24    // 全局变量；

func main() {
    var locali int = 42   // 局部变量；

    fmt.Println(globali,locali)
}

// 如果全局变量的首字母大写，那么它就是公开的全局变量，其他package也可以访问(使用)；
// 如果全局变量的首字母小写，那么它就是package内的全局变量，只有当前package内的代码可以访问；其他package的代码是不能访问的；
```

## 变量与常量：

* 常量：使用const关键字定义；

```go
常量可以是全局常量也可以是局部常量；

常量不可以修改，否则编译器会报错；

常量必须初始化，因为它无法二次赋值；

常量的大小写规则和变量是一致的；
```

```go
package main

import "fmt"

const globali int = 24

func main() {
    const locali int = 42
    fmt.Println(globali, locali)
}
```

## 指针类型：

Go延续使用了C语言的指针类型

```go
package main

import "fmt"

func main() {
    var value int = 42
    var pointer *int = &value
    fmt.Println(pointer,*pointer)
}

// & 取地址符；
// * 指针符号；
```

* 多级指针(了解)

```go
package main

import "fmt"

func main() {
    var value int = 42
    var p1 *int = &value
    var p2 **int = &p1
    var p3 ***int = &p2
    fmt.Println(p1,p2,p3)       // 三个不同的内存地址；
    fmt.Println(*p1,**p2,***p3)  // 42 42 42
}
```

```go
// 指针变量本质上就是一个整型变量，里面存储的值是另一个变量内存的地址。

// * 和 & 符号都只是它的语法糖，是用来在形式上方便使用和理解指针的。

// * 操作符存在两次内存读写，第一次获取指针变量的值，也就是内存地址，然后再去拿这个内存地址所在的变量内容。
```

如果普通的变量是一个储物箱，那么指针变量就是另一个储物箱，这个储物箱里存放了普通变量所在储物箱的钥匙。通过多级指针来读取变量值就好比在玩一个解密游戏。


## Go语言基础类型：

```go
package main

import "fmt"

func main() {
	// 有符号整数，可以表示正负
	var a int8 = 1  // 1字节
	var b int16 = 2 // 2字节
	var c int32 = 3 // 4字节
	var d int64 = 4 // 8字节
	fmt.Println(a, b, c, d)

	// 无符号整数，只能表示非负数
	var ua uint8 = 1
	var ub uint16 = 2
	var uc uint32 = 3
	var ud uint64 = 4
	fmt.Println(ua, ub, uc, ud)

	// int类型，在32位机器上占4个字节，在64位机器上占8个字节
	var e int = 5
	var ue uint = 5
	fmt.Println(e, ue)

	// bool 类型
	var f bool = true
	fmt.Println(f)

	// 字节类型
	var j byte = 'a'
	fmt.Println(j)

	// 字符串类型
	var g string = "abcdefg"
	fmt.Println(g)

	// 浮点数
	var h float32 = 3.14
	var i float64 = 3.141592653
	fmt.Println(h, i)
}

/*
不常用的数据类型:
	1.复数类型 complex64 和 complex128
	2.unicode字符类型 rune
	3.uintptr 指针类型

rune 和 byte 的关系就好比 Python 里面的 unicode 和 byte;
Java 语言里面的 char 和 byte 。
uintptr 相当于 C 语言里面的 void* 指针类型。
*/
```

