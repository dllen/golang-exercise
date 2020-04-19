package main

import (
	"fmt"
	"time"
)

//比较常见的一个场景是重试，第一个请求在指定超时时间内没有返回结果，这时重试第二次，取两次中最快返回的结果使用。
func main() {
	ret := make(chan string, 3)
	for i := 0; i < cap(ret); i++ {
		go call(ret)
	}
	select {
	case <-ret:
		// logic
	case <-time.After(3 * time.Second):
		// timeout
	}
}

func call(ret chan<- string) {
	// do something
	// ...
	fmt.Println("do something")
	ret <- "done"
}
