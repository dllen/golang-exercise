package main

import (
	"fmt"
	"net"
)

//网络 I/O
func main() {
	c, err := net.Dial("udp", "[fe80::1%lo0]:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
}
