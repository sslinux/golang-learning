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
