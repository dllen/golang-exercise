package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{1, 2, 4, 2, 4}
	var newS []int64
	// 做法是利用 reflect 直接替换数据指针
	// 但是这个不保证在以后的版本中一直可用 ╮(╯▽╰)╭
	*(*reflect.SliceHeader)(unsafe.Pointer(&newS)) = *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("type:%T value:%v", newS, newS)
}
