package main

import "fmt"

func main() {
	var s1 = "hello world"
	var b = []byte(s1) // 字符串转字节切片
	var s2 = string(b) // 字节切片转字符串；
	fmt.Println(b)
	fmt.Println(s2)
}
