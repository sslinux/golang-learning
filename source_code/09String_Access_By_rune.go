package main

import "fmt"

func main() {
	var s = "嘻哈China"
	for codepoint, runeValue := range s {
		fmt.Printf("%d %d ", codepoint, int32(runeValue))
	}

	a := s[:4]
	fmt.Println(a)
}
