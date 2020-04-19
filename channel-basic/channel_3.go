package main

import (
	"fmt"
	"time"
)

//限制最大并发数
func main() {
	// 最大并发数为 2
	limits := make(chan struct{}, 2)
	for i := 0; i < 10; i++ {
		go func() {
			// 缓冲区满了就会阻塞在这
			limits <- struct{}{}
			fmt.Println("do something")
			time.Sleep(time.Second)
			<-limits
		}()
	}
	time.Sleep(time.Second * 100)
}
