package main

import (
	"errors"
	"fmt"
)

func main() {
	//拆开一个被包装的 error
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	fmt.Println(errors.Unwrap(err3))
	fmt.Println(errors.Unwrap(errors.Unwrap(err3)))

	//判断被包装的 error 是否是包含指定错误
	//当多层调用返回的错误被一次次地包装起来，我们在调用链上游拿到的错误如何判断是否是底层的某个错误呢？
	//它递归调用 Unwrap 并判断每一层的 err 是否相等，如果有任何一层 err 和传入的目标错误相等，则返回 true。
	err1 = errors.New("new error")
	err2 = fmt.Errorf("err2: [%w]", err1)
	err3 = fmt.Errorf("err3: [%w]", err2)
	//Is As 两个方法已经预留了口子，可以由自定义的 error struct 实现并覆盖调用。
	fmt.Println(errors.Is(err2, err1))
	fmt.Println(errors.Is(err3, err2))
	fmt.Println(errors.Is(err3, err1))
}
