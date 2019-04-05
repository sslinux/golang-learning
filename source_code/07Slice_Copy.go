package main

import "fmt"

func main() {
	var s1 = make([]int, 5, 8)
	// 切片的访问和数组差不多
	for i := 0; i < len(s1); i++ {
		s1[i] = i + 1
	}
	var s2 = s1 // 浅拷贝

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// 尝试修改切片内容
	s2[0] = 255
	fmt.Println(s1)
	fmt.Println(s2)

	/*
		[1 2 3 4 5] 5 8
		[1 2 3 4 5] 5 8
		[255 2 3 4 5]
		[255 2 3 4 5]
		// 对s2[0]的修改，影响了s1[0]的值，说明s1和s2共享了底层数组；
	*/
}
