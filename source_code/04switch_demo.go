package main

import "fmt"

func main() {
	fmt.Println(prize1(60))
	fmt.Println(prize2(90))
}

func prize1(score int) string {
	// swithc的变量值匹配模式：
	switch score / 10 {
	case 0, 1, 2, 3, 4, 5:
		return "差"
	case 6, 7:
		return "及格"
	case 8:
		return "良"
	default:
		return "优"
	}
}

func prize2(score int) string {
	// switch的表达式匹配模式：
	switch {
	case score < 60:
		return "差"
	case score < 80:
		return "及格"
	case score < 90:
		return "良"
	default:
		return "优"
	}
}
