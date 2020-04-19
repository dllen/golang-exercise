package main

import (
	"fmt"
)

//Recover
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	panic(fmt.Errorf("Hello World"))
}
