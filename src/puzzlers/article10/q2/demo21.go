package main

func main() {
	// 示例1。
	ch1 := make(chan int, 1)
	ch1 <- 1
	// ch1 <- 2 // 通道已满，因此这里会造成阻塞。
	// 写入阻塞导致当前 goroutine 睡眠，又由于整个程序只有一个 goroutine，所以会报出 deadlock 错误
	// fatal error: all goroutines are asleep - deadlock!

	// 示例2。
	ch2 := make(chan int, 1)
	// elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞。fatal error: all goroutines are asleep - deadlock!
	// _, _ = elem, ok
	ch2 <- 1

	// 示例3。
	var ch3 chan int
	//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
	_ = ch3 // 这样编译器就不会提示未使用的变量错误了
}
