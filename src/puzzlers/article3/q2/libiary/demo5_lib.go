package lib
// 代码中指定的名字是其他地方引用时使用的名字。
// 代码库的文件路径是其他地方 import 的时候使用的名字

import "fmt"

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
