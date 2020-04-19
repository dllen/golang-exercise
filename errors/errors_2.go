package main

import (
	"errors"
	"fmt"
)

//自定义 struct
type WrapError struct {
	msg string
	err error
}

func (e *WrapError) Error() string {
	return e.msg
}

func (e *WrapError) Unwrap() error {
	return e.err
}

func main() {
	err := &WrapError{
		msg: "hhh",
		err: errors.New("hhh"),
	}
	fmt.Println(err)
}
