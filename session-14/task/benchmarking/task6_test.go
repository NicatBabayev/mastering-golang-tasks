package benchmarking

import "testing"

var arr []string

func BenchmarkConcatanateWithPlus(b *testing.B) {
	arr = createRandomArray2(10)
	for n := 0; n < b.N; n++ {
		concatanateWithPlus(arr)
	}
}
func BenchmarkConcatanateWithJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatanateWithJoin(arr)
	}
}
