package main

import "fmt"

func main() {
	var s = make([]int, 5, 8)
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	fmt.Println(s)

	var d = make([]int, 2, 6)
	var n = copy(d, s) // 返回的是2，n=2
	fmt.Println(n, d)
}
