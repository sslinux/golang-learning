# Golang流程控制

任何编程语言都要用的顺序、分支和循环；


程序 = 数据结构 + 算法

数据结构是内存数据关系的静态表示，算法是数据结构从一个状态变化到另一个状态需要执行的机器指令序列。

数据结构是静态的，算法是动态的。数据结构是状态，算法是状态的变化。


* 变量是数据结构的最小单位，分支与循环是算法逻辑的最小单位；


## if else语句

```go
package main

import "fmt"

func main() {
	fmt.Println(sign(max(min(24, 42), max(24, 42))))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}
// go语言没有三元操作符(a>b?a:b)
```

## switch语句：

`switch 有两种匹配模式：一种是变量值匹配，一种是表达式匹配。`

```go
package main

import "fmt"

func main() {
	fmt.Println(prize1(60))
	fmt.Println(prize2(90))
}

func prize1(score int) string {
	// swithc的变量值匹配模式：
	switch score / 10 {
	case 0, 1, 2, 3, 4, 5:
		return "差"
	case 6, 7:
		return "及格"
	case 8:
		return "良"
	default:
		return "优"
	}
}

func prize2(score int) string {
	// switch的表达式匹配模式：
	switch {
	case score < 60:
		return "差"
	case score < 80:
		return "及格"
	case score < 90:
		return "良"
	default:
		return "优"
	}
}
```

switch 还支持特殊的类型匹配语法，后续补充。


## for循环：

Go 语言虽然没有提供 while 和 do while 语句，不过这两个语句都可以使用 for 循环的形式来模拟。

我们平时使用 while 语句来写死循环 while(true) {}，Go 语言可以这么写

```go
package main

import "fmt"

func main() {
    for {
        fmt.Println("Golang is awesome")
    }
}
```

也可以这样写：

```go
package main

import "fmt"

func main() {
    for true {
        fmt.Println("Golang is awesome")
    }
}
```

* for 什么条件也不带的，相当于 loop {} 语句。
* for 带一个条件的相当于 while 语句。
* for 带三个语句的就是普通的 for 语句。

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Golang is Awesome")
    }
}
```

* 循环控制：

Go 语言支持 continue 和 break 语句来控制循环，这两个语句和其它语言没什么特殊的区别。

Go 语言还支持 goto 语句。

```go
计算机科学家迪杰.斯特拉很讨厌 goto 语句，他的那片广为流传的论文《Goto 有害》已经家喻户晓般地为人所知，可是他已经去世这么多年了，很遗憾还是没能阻挡 goto 语句死灰复燃般地继续存在。

我们作为普通用户平时还是尽量不要使用 goto 语句了，不论是谁来 review 你的代码只要是看到了 goto 语句肯定会发牢骚的。
```

