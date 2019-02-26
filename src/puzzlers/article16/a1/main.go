package main

import "fmt"

func main() {
	numNotify := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			numNotify <- struct{}{}
		}(i)

		<-numNotify
	}
}
