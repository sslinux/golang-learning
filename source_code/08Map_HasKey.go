package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	var score, ok = fruits["durin"] //此时字典中不存在该key
	if ok {
		fmt.Println(score)
	} else {
		fmt.Println("durin not exists")
	}

	fruits["durin"] = 0
	score, ok = fruits["durin"] // 此时字典中有这个key，但是值为0；
	if ok {
		fmt.Println(score)
	} else {
		fmt.Println("durin not exists")
	}
}
