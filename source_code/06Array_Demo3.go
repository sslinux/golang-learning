package main

import "fmt"

func main() {
	var squares [9]int
	for i := 0; i < len(squares); i++ {
		squares[i] = (i + 1) * (i + 1)
	}
	fmt.Println(squares)
}
