package main

import "fmt"

func main() {
	var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
