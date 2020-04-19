package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 20)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	// 当 c 被关闭后，取完里面的元素就会跳出循环
	for x := range c {
		fmt.Println(x)
	}
}
