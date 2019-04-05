package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}

	// 使用下标遍历：
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println("*******************************")

	// 使用for range 1:
	for index := range a {
		fmt.Println(index, a[index])
	}

	fmt.Println("*******************************")

	// 使用for range 2:
	for index, value := range a {
		fmt.Println(index, value)
	}
}
