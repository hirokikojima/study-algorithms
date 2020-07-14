package search

import (
	"testing"
)

func BenchmarkLinearSearch(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	search := NewSearch(list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		search.LinearSearch(6543)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	search := NewSearch(list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		search.BinarySearch(6543)
	}
}

func BenchmarkHashSearch(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	search := NewSearch(list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		search.HashSearch(6543)

	}
}