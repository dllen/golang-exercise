package main

import (
	"errors"
	"fmt"
)

//fmt.Errorf
//使用 %w 参数返回一个被包装的 error
func main() {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)
	fmt.Println(err3)
}
