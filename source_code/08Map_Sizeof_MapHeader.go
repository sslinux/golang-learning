package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m = map[string]int{
		"apple":  2,
		"pear":   3,
		"banana": 5,
	}
	fmt.Println(unsafe.Sizeof(m))
}
