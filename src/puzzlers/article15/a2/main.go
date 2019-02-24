package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	arr := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &arr) //0xc0000161e0

	sli := arr[:]
	fmt.Printf("%p\n", sli)     //0xc0000161e0
	fmt.Printf("%p\n", &sli[0]) //0xc0000161e0

	fmt.Printf("%p\n", &sli)          //0xc00000a080
	fmt.Println(unsafe.Pointer(&sli)) //0xc00000a080

	sliHeader := (*reflect.SliceHeader)(unsafe.Pointer(&sli))
	fmt.Printf("0x%10x\n", sliHeader.Data) //0xc0000161e0
}
