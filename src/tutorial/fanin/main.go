package main

import "fmt"

func gen(v string, times int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < times; i++ {
			ch <- v
		}
	}()
	return ch
}

func fanIn(times int, inputs []<-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < times; i++ {
			for _, input := range inputs {
				v := <-input
				ch <- v
			}
		}
	}()
	return ch
}

func main() {
	times := 10
	inputs := make([]<-chan string, 0, 3)
	for _, K := range []string{"A", "B", "C"} {
		inputs = append(inputs, gen(K, times))
	}
	for char := range fanIn(times, inputs) {
		fmt.Println(char)
	}
}
