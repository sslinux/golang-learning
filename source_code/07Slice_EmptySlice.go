package main

import "fmt"

func main() {
	var s1 []int
	var s2 []int = []int{}
	var s3 []int = make([]int, 0)

	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))
}
