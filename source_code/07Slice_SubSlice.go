package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	// start_index 和 end_index 不包含 end_index
	// [start_index:end_index]
	var s2 = s1[2:5]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}
