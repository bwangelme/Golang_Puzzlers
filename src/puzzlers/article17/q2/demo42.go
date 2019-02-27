package main

import "fmt"

func main() {
	// 示例1。
	// value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// switch 1 + 3 { // 这条语句无法编译通过。
	// case value1[0], value1[1]:
	// 	fmt.Println("0 or 1")
	// case value1[2], value1[3]:
	// 	fmt.Println("2 or 3")
	// case value1[4], value1[5], value1[6]:
	// 	fmt.Println("4 or 5 or 6")
	// }

	// 示例2。
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	// puzzlers/article17/q2/demo42.go:22:10: constant 3.6 truncated to integer
	case 2, 3.6:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}
