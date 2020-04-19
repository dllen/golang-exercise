package main

import (
	"fmt"
	"reflect"
)

//通过 kind（）处理不同分支
func main() {
	a := 1
	t := reflect.TypeOf(a)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	}
}
