package main

import (
	"fmt"
	"time"
)

//利用 close 广播
func main() {
	c := make(chan struct{})
	for i := 0; i < 5; i++ {
		go do(c)
	}
	close(c)
	time.Sleep(time.Second * 10)
}

func do(c <-chan struct{}) {
	// 会阻塞直到收到 close
	<-c
	fmt.Println("hello")
}
