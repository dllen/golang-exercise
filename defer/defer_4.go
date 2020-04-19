package main

import (
	"fmt"
	"os"
)

//channel 关闭
func main() {
	fd, _ := os.Open("./defer.txt")
	errc := make(chan error, 1)
	// 主动关闭，减小 GC 压力。
	defer close(errc)

	var buf [1]byte
	n, err := fd.Read(buf[:1])
	if n == 0 || err != nil {
		errc <- fmt.Errorf("read byte = %d, err = %v", n, err)
	}
}
