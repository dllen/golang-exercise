package test

/*
无并行
GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
有并行
GOGC=off go test -cpu 4 -run none -bench . -benchtime 3s
*/

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}

func addConcurrent(goroutines int, numbers []int) int {
	var v int64
	totalNumbers := len(numbers)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoroutine {
				end = totalNumbers
			}

			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}

			atomic.AddInt64(&v, int64(lv))
			wg.Done()
		}(g)
	}

	wg.Wait()

	return int(v)
}

func BenchmarkSequential(b *testing.B) {
	var numbers []int
	for i := 0; i < 100000; i++ {
		numbers = append(numbers, rand.Intn(100000))
	}
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	var numbers []int
	for i := 0; i < 100000; i++ {
		numbers = append(numbers, rand.Intn(100000))
	}
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}
