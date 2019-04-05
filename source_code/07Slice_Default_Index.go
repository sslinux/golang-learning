package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	var s2 = s1[:5]
	var s3 = s1[3:]
	var s4 = s1[:]

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))

	/*
	   -----------------------
	   [1 2 3 4 5 6 7] 7 7
	   [1 2 3 4 5] 5 7
	   [4 5 6 7] 4 4
	   [1 2 3 4 5 6 7] 7 7
	*/
}
