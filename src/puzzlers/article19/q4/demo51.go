package main

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
)

func main() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer func() {
		fmt.Println("last defer")
		panic(errors.New("panic in defer"))
	}()
}
