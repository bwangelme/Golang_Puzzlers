package main

import (
	"fmt"
)

func main() {
	signal := make(chan struct{})

	defer func() {
		// Go 语言运行时系统自行抛出的panic，例如 deadlock 是无法被恢复的
		if p := recover(); p != nil {
			fmt.Printf("Panic: %v\n", p)
		}
	}()

	go func() {
		<-signal
	}()

	<-signal
}
