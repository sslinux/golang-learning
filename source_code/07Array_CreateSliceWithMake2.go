package main

import "fmt"

func main() {
	var s1 = make([]int, 5, 8)
	s2 := make([]int, 8)
	fmt.Println(s1)
	fmt.Println(s2)
}
