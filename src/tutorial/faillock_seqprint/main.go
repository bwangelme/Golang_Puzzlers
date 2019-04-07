package main

import (
	"fmt"
	"log"
	"sync"
)

type FailLock struct {
	mu        *sync.Mutex
	cond      *sync.Cond
	holdCount int
}

func NewFailLock() sync.Locker {
	mu := new(sync.Mutex)
	cond := sync.NewCond(mu)

	return &FailLock{
		holdCount: 0,
		mu:        mu,
		cond:      cond,
	}
}

func (fl *FailLock) Lock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	fl.holdCount++
	if fl.holdCount == 1 {
		return
	}

	fl.cond.Wait()
}

func (fl *FailLock) Unlock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	if fl.holdCount == 0 {
		log.Fatal("unlock of UnLocked mutex")
	}

	fl.holdCount--
	if fl.holdCount != 0 {
		fl.cond.Signal()
	}
}

var (
	end = make(chan struct{})
	i   int
)

func threadPrint(threadNum int, threadName string, mu sync.Locker) {
	for i < 30 {
		mu.Lock()
		if i >= 30 {
			mu.Unlock()
			continue
		}
		if i < 3 && i%3 != threadNum {
			mu.Unlock()
			continue
		}

		fmt.Printf("%d: %s\n", i, threadName)
		i += 1
		mu.Unlock()
	}
	end <- struct{}{}
}

func main() {
	mu := NewFailLock()
	names := []string{"A", "B", "C"}

	for idx, name := range names {
		go threadPrint(idx, name, mu)
	}

	for _ = range names {
		<-end
	}
}
