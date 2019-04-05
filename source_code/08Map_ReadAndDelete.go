package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	//读取元素：
	var score = fruits["banana"]
	fmt.Println(score)

	// 增加或修改元素：
	fruits["pear"] = 3
	fmt.Println(fruits)

	// 删除元素：
	delete(fruits, "pear")
	fmt.Println(fruits)
}
