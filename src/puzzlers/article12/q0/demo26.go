package main

import "log"

type Printer func(string) (int, error)

func printToStd(contents string) (bytesNum int, err error) {
	log.Println(contents)
	return len(contents), nil
}

func main() {
	var p Printer
	p = printToStd
	p("something")
}
