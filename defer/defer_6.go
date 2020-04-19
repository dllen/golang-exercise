package main

import (
	"fmt"
	"time"
)

func main() {
	do()
}

func do() {
	defer log("do")()

	// ... 一些逻辑

	time.Sleep(1 * time.Second)
}

//记日志
func log(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() { fmt.Printf("exit %s (%s)", msg, time.Since(start)) }
}
