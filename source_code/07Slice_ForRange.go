package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index := range s {
		fmt.Println(index, s[index])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
