package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 示例1。
	// 只能发不能收的通道。
	var uselessChan = make(chan<- int, 1)
	// 只能收不能发的通道。
	var anotherUselessChan = make(<-chan int, 1)
	// 这里打印的是可以分别代表两个通道的指针的16进制表示。
	fmt.Printf("The useless channels: %v, %v\n", uselessChan, anotherUselessChan)

	// 示例2。
	intChan1 := make(chan int, 3)
	SendInt(intChan1)
	fmt.Println(<-intChan1)

	// 示例4。
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	// 示例5。
	_ = GetIntChan(getIntChan)
}

// 示例2。
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

// 示例3。
type Notifier interface {
	SendInt(ch chan<- int)
}

// 示例4。
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= num; i++ {
		ch <- r.Intn(i * 10)
	}
	close(ch)
	return ch
}

// 示例5。
type GetIntChan func() <-chan int
