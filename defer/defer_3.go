package main

import (
	"fmt"
	"os"
)

//文件 I/O
func main() {
	f, err := os.Open("./defer.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
}
