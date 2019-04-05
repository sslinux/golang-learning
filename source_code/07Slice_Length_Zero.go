package main

import "fmt"

func main() {
	var s1 []int
	var s2 = []int{}
	var s3 = make([]int, 0)

	// new函数返回的是指针类型，所以需要使用 * 号来解引用；
	var s4 = *new([]int)

	fmt.Println(len(s1), len(s2), len(s3), len(s4))
	fmt.Println(cap(s1), cap(s2), cap(s3), cap(s4))
	fmt.Println(s1, s2, s3, s4)
}

/*
--------------------
0 0 0 0
0 0 0 0
[] [] [] []
*/
