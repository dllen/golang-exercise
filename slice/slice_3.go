package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int64{1, 2, 4, 2, 4}
	fmt.Printf("s1: %v\n", RemoveDuplicates(s1))
	fmt.Printf("s2: %v\n", RemoveDuplicates2(s1))
	fmt.Printf("s3: %v\n", RemoveDuplicates3(s1))
}

func RemoveDuplicates(s []int64) []int64 {
	var ret []int64
	for _, v := range s {
		found := false
		for _, v2 := range ret {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, v)
		}
	}
	return ret
}

func RemoveDuplicates2(s []int64) []int64 {
	ret := s[:0]
	// 利用 struct{}{} 减少内存占用
	assist := map[int64]struct{}{}
	for _, v := range s {
		if _, ok := assist[v]; !ok {
			assist[v] = struct{}{}
			ret = append(ret, v)
		}
	}
	return ret
}

func RemoveDuplicates3(a []int64) []int64 {
	sort.Ints(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		in[j] = in[i]
	}
	result := in[:j+1]
	return result
}
