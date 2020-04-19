package main

import (
	"fmt"
)

func main() {
	s1 := []int64{1, 2, 4, 2, 4}
	fmt.Printf("s1: %v\n", SliceChunk(s1, 2))
}

func SliceChunk(s []int64, size int) [][]int64 {
	var ret [][]int64
	for size < len(s) {
		// s[:size:size] 表示 len 为 size，cap 也为 size，第二个冒号后的 size 表示 cap
		s, ret = s[size:], append(ret, s[:size:size])
	}
	ret = append(ret, s)
	return ret
}
