package main

import (
	"fmt"
	"strconv"
)

var container = []string{"zero", "one", "two"}

func main() {
	// container := map[int]string{0: "zero", 1: "one", 2: "two"}

	// 方式1。
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)
	// panic: interface conversion: interface {} is []int, not []string

	// goroutine 1 [running]:
	// main.main()
	//        /Users/michaeltsui/Github/Golang/geektime/Golang_Puzzlers/src/puzzlers/article6/q1/demo12.go:16 +0x45
	// exit status 2
	// _ = interface{}([]int{2, 4, 5}).([]string)
	// 如果在进行类型断言的判断的时候没有加 ok 返回值，那么判断失败的时候，程序就会抛出 panic
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n", container[1], container)

	// 方式2。
	elem, err := getElement([]int{1, 2, 3})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n", elem, container)
}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	case []int:
		elem = strconv.Itoa(t[1])
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return
	}
	return
}
