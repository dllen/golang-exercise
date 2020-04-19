package main

import (
	"errors"
	"fmt"
	"reflect"
)

type T struct{}

//返回值也是 Value 类型，对于错误，可以转为 interface 之后断言
func main() {
	name := "Do"
	t := &T{}
	ret := reflect.ValueOf(t).MethodByName(name).Call(nil)
	fmt.Printf("strValue: %[1]v\nerrValue: %[2]v\nstrType: %[1]T\nerrType: %[2]T", ret[0], ret[1].Interface().(error))
}

func (t *T) Do() (string, error) {
	return "hello", errors.New("new error")
}
