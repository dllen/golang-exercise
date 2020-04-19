package popcnt

import (
	"testing"
)

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

//因为函数是内联的，所以编译器现在可以看到它没有副作用。  popcnt不会影响任何全局变量的状态。 这样，调用就被消除了。
// func BenchmarkPopcnt(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		popcnt(uint64(i))
// 	}
// }

var Result uint64

func BenchmarkPopcnt(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = popcnt(uint64(i))
	}
	Result = r
}
