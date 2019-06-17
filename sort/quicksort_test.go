package sort

import (
	"math/rand"
	"testing"
)

const (
	maxIntValue = 10000
)

func TestLomutoQuickSort(t *testing.T) {
	arr := []int{1,4,7,2,3,9,8,1}
	LomutoQuickSort(arr)
	for i := 0; i < len(arr) - 2; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("Array %v is not sorted!", arr)
		}
	}
}

func TestHoareQuickSort(t *testing.T) {
	arr := []int{1,4,7,2,3,9,8,0,5,2}
	HoareQuickSort(arr)
	for i := 0; i < len(arr) - 2; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("Array %v is not sorted!", arr)
		}
	}
}

func generateRandomArray(length int) []int {
	var a []int
	for i := 0; i < length; i++ {
		a = append(a, rand.Intn(maxIntValue))
	}
	return a
}

func benchmarkQuickSort(length int, b *testing.B) {
	a := generateRandomArray(length)
	b.ResetTimer()
	LomutoQuickSort(a)
}

func BenchmarkQuickSort10(b *testing.B) {
	benchmarkQuickSort(10, b)
}

func BenchmarkQuickSort100(b *testing.B) {
	benchmarkQuickSort(100, b)
}

func BenchmarkQuickSort1000(b *testing.B) {
	benchmarkQuickSort(1000, b)
}

func BenchmarkQuickSort10000(b *testing.B) {
	benchmarkQuickSort(10000, b)
}

func BenchmarkQuickSort100000(b *testing.B) {
	benchmarkQuickSort(100000, b)
}

func BenchmarkQuickSort1000000(b *testing.B) {
	benchmarkQuickSort(1000000, b)
}

func BenchmarkQuickSort10000000(b *testing.B) {
	benchmarkQuickSort(10000000, b)
}