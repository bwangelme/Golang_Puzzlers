package main

import . "os"
import "fmt"


// # puzzlers/article5/a1
// puzzlers/article5/a1/main.go:3:8: imported and not used: "os"
// 注意，这里将 OS 中的常量 PathSeparator 重声明了，但是报错不是 redeclared

var PathSeparator = '\\'

func main() {
    fmt.Printf("PathSeparator %s \n", string(PathSeparator))
}