package main

import (
	"fmt"
	"sync/atomic"
)

var (
	end = make(chan struct{})
	i   int32
)

func threadPrint(threadNum int32, threadName string) {

	for {
		v := atomic.LoadInt32((*int32)(&i))
		fmt.Println(v)
		if v >= 50 {
			break
		}

		if v%5 == threadNum {
			fmt.Printf("%d: %s\n", i, threadName)
			atomic.AddInt32((*int32)(&i), int32(1))
		} else {
			//runtime.Gosched()
			continue
		}
	}
	end <- struct{}{}
}

func main() {
	names := []string{"A", "B", "C", "D", "E"}

	for idx, name := range names {
		go threadPrint(int32(idx), name)
	}

	for _ = range names {
		<-end
	}
}
