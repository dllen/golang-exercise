package main

import (
	"errors"
	"fmt"
)

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

//提取指定类型的错误
//Is As 两个方法已经预留了口子，可以由自定义的 error struct 实现并覆盖调用。
func main() {
	var targetErr *ErrorString
	err := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})

	//这个和上面的 errors.Is 大体上是一样的，区别在于 Is 是严格判断相等，即两个 error 是否相等。
	//而 As 则是判断类型是否相同，并提取第一个符合目标类型的错误，用来统一处理某一类错误。
	fmt.Println(errors.As(err, &targetErr))
}
