package main

import (
	"fmt"
)

func main() {
	s1 := []int64{1, 2, 4}
	s2 := FilterSlice(s1, func(x int64) bool { return x%2 == 0 })
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func FilterSlice(s []int64, filter func(x int64) bool) []int64 {
	// 返回的新切片
	// s[:0] 这种写法是创建了一个 len 为 0，cap 为 len(s) 即和原始切片最大容量一致的切片
	// 因为是过滤，所以新切片的元素总个数一定不大于比原始切片，这样做减少了切片扩容带来的影响
	// 同时，也有一个问题，因为 newS 和 s 共享底层数组，那么过滤后 s 也会被修改！
	newS := s[:0]
	// 遍历，对每个元素执行 filter，符合条件的加入新切片中
	for _, x := range s {
		if !filter(x) {
			newS = append(newS, x)
		}
	}
	return newS
}
