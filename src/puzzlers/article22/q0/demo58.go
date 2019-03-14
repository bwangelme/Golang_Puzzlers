package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

// protecting 用于指示是否使用互斥锁来保护数据写入。
// 若值等于0则表示不使用，若值大于0则表示使用。
// 改变该变量的值，然后多运行几次程序，并观察程序打印的内容。
var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 1,
		"It indicates whether to use a mutex to protect data writing.")
}

// Buf for string
type Buf struct {
	buffer   bytes.Buffer
	mu       sync.Mutex
	isLocked bool
}

// WriteLines func
func (buf *Buf) WriteLines(lines []string) (int, error) {
	// 不上锁的情况下，后写入的数据会将之前写入的覆盖掉
	if buf.isLocked {
		buf.mu.Lock()
		defer buf.mu.Unlock()
	}

	total := 0
	for _, line := range lines {
		n, err := buf.buffer.Write([]byte(line))
		if err != nil {
			return total, err
		}
		total += n
	}

	return total, nil
}

// ReadAll Read all data in buffer
func (buf *Buf) ReadAll() ([]byte, error) {
	if buf.isLocked {
		buf.mu.Lock()
		defer buf.mu.Unlock()
	}

	data, err := ioutil.ReadAll(&buf.buffer)
	return data, err
}

// NewBuf function
func NewBuf(isLocked bool) *Buf {
	var buf Buf

	buf.isLocked = isLocked

	return &buf
}

func main() {
	flag.Parse()

	const (
		max1 = 5  // 代表启用的goroutine的数量。
		max2 = 10 // 代表每个goroutine需要写入的数据块的数量。
		max3 = 3  // 代表每个数据块中需要有多少个重复的数字。
	)

	var buf = NewBuf(bool(protecting > 0))
	// sign 代表信号的通道。
	sign := make(chan struct{}, max1)

	for i := 1; i <= max1; i++ {
		go func(id int, buf *Buf) {
			defer func() {
				sign <- struct{}{}
			}()
			var lines []string
			for j := 1; j <= max2; j++ {
				// 准备数据。
				// 这里准备的是一行的数据
				header := fmt.Sprintf("[id: %d, iteration: %d]",
					id, j)
				body := ""
				for k := 0; k < max3; k++ {
					body += fmt.Sprintf(" %d", id*j)
				}
				data := header + body + "\n"
				lines = append(lines, data)
			}
			// 写入数据。
			_, err := buf.WriteLines(lines)
			if err != nil {
				log.Printf("error: %s [%d]", err, id)
			}
		}(i, buf)
	}

	for i := 0; i < max1; i++ {
		<-sign
	}
	data, err := buf.ReadAll()
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents:\n%s", data)
}
