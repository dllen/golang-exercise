package test

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkSortSequential(b *testing.B) {
	var numbers []int
	for i := 0; i < 100000; i++ {
		numbers = append(numbers, rand.Intn(100000))
	}
	for i := 0; i < b.N; i++ {
		bubbleSort(numbers)
	}
}

func BenchmarkSortConcurrent(b *testing.B) {
	var numbers []int
	for i := 0; i < 100000; i++ {
		numbers = append(numbers, rand.Intn(100000))
	}
	for i := 0; i < b.N; i++ {
		bubbleSortConcurrent(runtime.NumCPU(), numbers)
	}
}

func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func bubbleSortConcurrent(goroutines int, numbers []int) {
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

			bubbleSort(numbers[start:end])
			wg.Done()
		}(g)
	}

	wg.Wait()

	// Ugh, we have to sort the entire list again.
	bubbleSort(numbers)
}

func sweep(numbers []int, currentPass int) bool {
	var idx int
	idxNext := idx + 1
	n := len(numbers)
	var swap bool

	for idxNext < (n - currentPass) {
		a := numbers[idx]
		b := numbers[idxNext]
		if a > b {
			numbers[idx] = b
			numbers[idxNext] = a
			swap = true
		}
		idx++
		idxNext = idx + 1
	}
	return swap
}
