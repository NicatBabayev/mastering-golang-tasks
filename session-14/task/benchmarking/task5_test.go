package benchmarking

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RunTask5()
	}
}
