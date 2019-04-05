package main

import "fmt"

func main() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b = a[2:6]
	fmt.Println(b)
	a[4] = 100
	fmt.Println(b)
}
