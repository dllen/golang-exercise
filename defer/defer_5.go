package main

import (
	"fmt"
	"sync"
	"time"
)

type A struct {
	t int
	sync.Mutex
}

//避免死锁
func main() {
	a := new(A)
	for i := 0; i < 2000; i++ {
		go a.incr()
	}
	// 此处用 sleep 简单模拟等待同步，实际这样写不严谨，可用 waitGroup、channel 等
	time.Sleep(500 * time.Millisecond)
	fmt.Println(a.t)
}

func (a *A) incr() {
	a.Lock()
	defer a.Unlock()

	// 模拟 ... 一堆逻辑

	// 然后 ... 中间有好几个 return 出口

	// 如果我们不用 defer，就要在每个 return 都写上 a.Unlock，不然就可能会造成死锁
	a.t++
}
