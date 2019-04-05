package main

// 只申明类型，不赋初值。这时编译器会给数组默认赋上「零值」。
// 数组的零值就是所有内部元素的零值。

import "fmt"

func main() {
	var a [9]int
	fmt.Println(a)
}
