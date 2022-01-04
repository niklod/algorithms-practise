package sorting

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSortInt_1000(b *testing.B) {
	data := generateIntSlice(b, 1000)
	b.ResetTimer()
	BubbleSortInt(data)
}

func BenchmarkBubbleSortInt_100000(b *testing.B) {
	data := generateIntSlice(b, 100000)
	b.ResetTimer()
	BubbleSortInt(data)
}

func BenchmarkBubbleSortInt_1000000(b *testing.B) {
	data := generateIntSlice(b, 1000000)
	b.ResetTimer()
	BubbleSortInt(data)
}

func BenchmarkBubbleSortInt_100000000(b *testing.B) {
	data := generateIntSlice(b, 100000000)
	b.ResetTimer()
	BubbleSortInt(data)
}

func BenchmarkBubbleSortInt_1000000000(b *testing.B) {
	data := generateIntSlice(b, 1000000000)
	b.ResetTimer()
	BubbleSortInt(data)
}

func generateIntSlice(b *testing.B, len int) []int {
	b.Helper()
	res := make([]int, len)

	for i := 0; i < len; i++ {
		res[i] = rand.Intn(1000000)
	}

	return res
}
