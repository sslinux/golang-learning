package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5}
	fmt.Println(s1, len(s1), cap(s1))

	// 对满容的切片进行追加会分离底层数组
	var s2 = append(s1, 6)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// 对非满容的切片进行追加会共享底层数组
	var s3 = append(s2, 7)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	/*
		运行结果：
		[1 2 3 4 5] 5 5

		[1 2 3 4 5] 5 5
		[1 2 3 4 5 6] 6 10  // s2在s1的基础上进行扩容，容量翻倍，此时s2是 非满容切片；

		[1 2 3 4 5 6] 6 10
		[1 2 3 4 5 6 7] 7 10  // s3在s2的基础上append元素，只会增加length，因为s2非满容；
	*/
}
