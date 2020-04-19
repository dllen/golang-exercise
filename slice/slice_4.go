package main

import (
	"fmt"
)

func main() {
	s1 := []int64{1, 2, 4, 2, 4}
	fmt.Printf("s1: %v\n", Reversing(s1))
	fmt.Printf("s2: %v\n", Reversing2(s1))
}

func Reversing(s []int64) []int64 {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
	return s
}

func Reversing2([]int64) []int64 {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}
