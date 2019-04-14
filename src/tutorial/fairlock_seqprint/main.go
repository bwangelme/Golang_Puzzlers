package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const endNum = 9

type FairLock struct {
	mu        *sync.Mutex
	cond      *sync.Cond
	holdCount int
}

func NewFairLock() sync.Locker {
	mu := new(sync.Mutex)
	cond := sync.NewCond(mu)

	return &FairLock{
		holdCount: 0,
		mu:        mu,
		cond:      cond,
	}
}

func (fl *FairLock) Lock() {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	fl.holdCount++
	if fl.holdCount == 1 {
		return
	}

	fl.cond.Wait()
}

func (fl *FairLock) Unlock() {
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
	for i < endNum {
		fmt.Println("......", threadNum, threadName, "Try Locked")
		mu.Lock()
		if i >= endNum {
			mu.Unlock()
			continue
		}
		if i < 3 && i%3 != threadNum {
			fmt.Println("......", threadNum, threadName, "Unlock")
			mu.Unlock()
			continue
		}

		fmt.Printf("%d: %s\n", i, threadName)
		i += 1
		fmt.Println("......", threadNum, threadName, "Worked")
		mu.Unlock()
	}
	end <- struct{}{}
}

func main() {
	mu := NewFairLock()
	names := []string{"A", "B", "C"}

	for idx, name := range names {
		time.Sleep(time.Second * 1)
		go threadPrint(idx, name, mu)
	}

	for _ = range names {
		<-end
	}
}
