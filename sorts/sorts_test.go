package sorts

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const size int = 10000

var list = make([]int, size)
var prng = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fillIntSlice(b)
		BubbleSort(list)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fillIntSlice(b)
		InsertionSort(list)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fillIntSlice(b)
		SelectionSort(list)
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fillIntSlice(b)
		sort.Sort(sort.IntSlice(list))
	}
}

func fillIntSlice(b *testing.B) {
	b.StopTimer()
	for i := range list {
		list[i] = prng.Int()
	}
	b.StartTimer()
}