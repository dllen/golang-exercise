package main

import (
	"fmt"
	"reflect"
)

type IT interface {
	test1()
}

type T struct {
	A string
}

func (t *T) test1() {}

//判断实例是否实现了某接口
func main() {
	t := &T{}
	ITF := reflect.TypeOf((*IT)(nil)).Elem()
	tv := reflect.TypeOf(t)
	fmt.Println(tv.Implements(ITF))
}
