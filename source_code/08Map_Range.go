package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	for key, val := range fruits {
		fmt.Println(key, val)
	}

	for key := range fruits {
		fmt.Println(key)
	}
}
