package main

import (
	"fmt"
	"time"
)

// 利用 time.After 实现
func main() {
	done := do()
	select {
	case <-done:
		// logic
	case <-time.After(3 * time.Second):
		// timeout
	}
}

func do() <-chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		// do something
		// ...
		fmt.Println("do something")
		done <- struct{}{}
	}()
	return done
}
