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
		if v >= 30 {
			break
		}

		if v%3 == threadNum {
			fmt.Printf("%d: %s\n", i, threadName)
			atomic.AddInt32((*int32)(&i), int32(1))
		} else {
			continue
		}
	}
	end <- struct{}{}
}

func main() {
	names := []string{"A", "B", "C"}

	for idx, name := range names {
		go threadPrint(int32(idx), name)
	}

	for _ = range names {
		<-end
	}
}
