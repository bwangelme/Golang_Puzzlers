package main

import "fmt"

func main() {
	ch1 := make(chan int)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			// 第一个写入的被立刻接收走了，第二个会阻塞，但也会先打印提示信息。
			// 所以打印的结果是这样的:
			// Sender: sending element 0...
			// Sender: sending element 1...
			// Receiver: received an element: 0
			// Receiver: received an element: 1
			// Sender: sending element 2...
			// Sender: sending element 3...
			// Receiver: received an element: 2
			// Receiver: received an element: 3
			// ......
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
