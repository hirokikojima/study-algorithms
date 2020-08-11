package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

const N = 100000

func TestSelectionSort(t *testing.T) {
	expect := []int{0,1,2,3,4,5,6,7,8,9}
	list := []int{0,9,8,7,6,5,4,3,2,1}

	sort := NewSort(list)

	sort.SelectionSort()

	if !reflect.DeepEqual(list, expect) {
		t.Error(list)
	}
}

func TestBubbleSort(t *testing.T) {
	expect := []int{0,1,2,3,4,5,6,7,8,9}
	list := []int{0,9,8,7,6,5,4,3,2,1}

	sort := NewSort(list)

	sort.BubbleSort()

	if !reflect.DeepEqual(list, expect) {
		t.Error(list)
	}
}

func TestInsertionSort(t *testing.T) {
	expect := []int{0,1,2,3,4,5,6,7,8,9}
	list := []int{0,9,8,7,6,5,4,3,2,1}

	sort := NewSort(list)

	sort.InsertionSort()

	if !reflect.DeepEqual(list, expect) {
		t.Error(list)
	}
}

func TestQuickSort(t *testing.T) {
	expect := []int{0,1,2,3,4,5,6,7,8,9}
	list := []int{0,9,8,7,6,5,4,3,2,1}

	sort := NewSort(list)

	sort.QuickSort()

	if !reflect.DeepEqual(list, expect) {
		t.Error(list)
	}
}

func TestHeapSort(t *testing.T) {
	expect := []int{9,8,7,6,5,4,3,2,1,0}
	list := []int{0,9,8,7,6,5,4,3,2,1}

	sort := NewSort(list)

	sort.HeapSort()

	if !reflect.DeepEqual(list, expect) {
		t.Error(list)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	list := []int{}
	for i := 0; i < N; i++ {
		list = append(list, rand.Intn(N))
	}

	sort := NewSort(list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.SelectionSort()
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	list := []int{}
	for i := 0; i < N; i++ {
		list = append(list, rand.Intn(N))
	}

	sort := NewSort(list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.BubbleSort()
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	list := []int{}
	for i := 0; i < N; i++ {
		list = append(list, rand.Intn(N))
	}

	sort := NewSort(list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.InsertionSort()
	}
}

// func BenchmarkQuickSort(b *testing.B) {
// 	list := []int{}
// 	for i := 0; i < N; i++ {
// 		list = append(list, rand.Intn(N))
// 	}

// 	sort := NewSort(list)

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		sort.QuickSort()
// 	}
// }

func BenchmarkHeapSort(b *testing.B) {
	list := []int{}
	for i := 0; i < N; i++ {
		list = append(list, rand.Intn(N))
	}

	sort := NewSort(list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.HeapSort()
	}
}