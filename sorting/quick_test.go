package sorting_test

import (
	"testing"

	"github.com/niklod/algorithms-practise/sorting"
)

func BenchmarkQuickSortIntSortInt_1000(b *testing.B) {
	data := generateIntSlice(b, 1000)
	b.ResetTimer()
	sorting.QuickSortInt(data)
}

func BenchmarkQuickSortIntSortInt_100000(b *testing.B) {
	data := generateIntSlice(b, 100000)
	b.ResetTimer()
	sorting.QuickSortInt(data)
}

func BenchmarkQuickSortIntSortInt_1000000(b *testing.B) {
	data := generateIntSlice(b, 1000000)
	b.ResetTimer()
	sorting.QuickSortInt(data)
}

func BenchmarkQuickSortIntSortInt_100000000(b *testing.B) {
	data := generateIntSlice(b, 100000000)
	b.ResetTimer()
	sorting.QuickSortInt(data)
}

// func BenchmarkQuickSortIntSortInt_1000000000(b *testing.B) {
// 	data := generateIntSlice(b, 1000000000)
// 	b.ResetTimer()
// 	sorting.QuickSortInt(data)
// }
