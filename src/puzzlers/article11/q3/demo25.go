package main

import (
	"fmt"
)

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var numbers = []int{1, 2, 3}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The first candidate case is selected.")
	case getChan(1) <- getNumber(1):
		// 因为第二个 channel 的 buf 为0，所以写入值后这个 channel 也是阻塞的，所以这个 case 分支也不会被触发
		fmt.Println("The second candidate case is selected.")
	case getChan(2) <- getNumber(2):
		fmt.Println("The third candidate case is selected")
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]: %d\n", i, numbers[i])
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
