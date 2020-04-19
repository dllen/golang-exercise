package main

import (
	"fmt"
	"reflect"
)

type T struct{}

//动态调用函数（无参数）
func main() {
	name := "Do"
	t := &T{}
	reflect.ValueOf(t).MethodByName(name).Call(nil)
}

func (t *T) Do() {
	fmt.Println("hello")
}
