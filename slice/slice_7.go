package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s1 := []int64{1, 2, 4, 2, 4}
	fmt.Printf("s1: %v\n", Shuffling(s1))
}

//Fisher–Yates algorithm
func Shuffling(a []int64) []int64 {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}
