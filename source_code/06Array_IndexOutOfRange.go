package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = 101
	a[b] = 255
	fmt.Println(a)
}
