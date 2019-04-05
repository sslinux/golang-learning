package main

import "fmt"

func main() {
	var s1 []int = make([]int, 5, 8)
	var s2 []int = make([]int, 8) //满容切片:length和capacity
	fmt.Println(s1)
	fmt.Println(s2)
}
