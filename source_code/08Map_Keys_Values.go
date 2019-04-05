package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	var names = make([]string, 0, len(fruits))
	var scores = make([]int, 0, len(fruits))
	// var names = make([]string, len(fruits))  // 已经是满容切片，需要扩展；
	// var scores = make([]int, len(fruits))

	for name, score := range fruits {
		names = append(names, name)
		scores = append(scores, score)
	}

	fmt.Println(names, scores)
}
