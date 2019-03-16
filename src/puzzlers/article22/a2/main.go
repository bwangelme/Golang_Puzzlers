package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex

	lock.Lock()
	fmt.Printf("Procted Data")
	// 这里解锁的是另外一个锁，所以 lock 再次锁定的时候，就会提示死锁错误
	// 可以传递 *sync.Muxtex 类型，这样就可以解锁当前锁了
	Unlock(lock)
	lock.Lock()
	lock.Unlock()

	// Out
	// Procted Datafatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [semacquire]:
	// sync.runtime_SemacquireMutex(0xc0000160a4, 0xc000016000)
	// 		/usr/local/opt/go/libexec/src/runtime/sema.go:71 +0x3d
	// sync.(*Mutex).Lock(0xc0000160a0)
	// 		/usr/local/opt/go/libexec/src/sync/mutex.go:134 +0xff
	// main.main()
	// 		/Users/michaeltsui/Github/Golang/geektime/Golang_Puzzlers/src/puzzlers/article22/a2/main.go:14 +0x8d
	// exit status 2
}

// Unlock unlock something
func Unlock(lock sync.Mutex) {
	lock.Unlock()
}
