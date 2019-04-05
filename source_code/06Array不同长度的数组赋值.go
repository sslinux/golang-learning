package main

import "fmt"

func main() {
	var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b [10]int

	b = a
	fmt.Println(b)
}
