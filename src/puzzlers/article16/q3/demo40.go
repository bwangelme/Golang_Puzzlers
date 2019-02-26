package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fmt.Println("Trigger Num is", i)
				fn()
				atomic.AddUint32(&count, 2)
				break
			}
			// time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 5; i++ {
		go func(i uint32) {
			// 闭包捕获的是函数定义时其外部函数的变量，例如这里捕获的是i，而不是i*2
			fn := func() {
				fmt.Println("Catch the num", i)
			}
			trigger(i*2, fn)
		}(i)
	}
	trigger(10, func() {})
}

// Out
// Trigger Num is 0
// Catch the num 0
// Trigger Num is 2
// Catch the num 1
// Trigger Num is 4
// Catch the num 2
// Trigger Num is 6
// Catch the num 3
// Trigger Num is 8
// Catch the num 4
// Trigger Num is 10
