package main

import (
	"slices"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, len(randomSlice))
		copy(slice, randomSlice)
		QuickSort(slice, sliceSize, 0, len(slice)-1)
	}
}

func BenchmarkQuickSortSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, len(sortedSlice))
		copy(slice, sortedSlice)
		QuickSort(slice, sliceSize, 0, len(slice)-1)
	}
}

func BenchmarkPDQSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, len(randomSlice))
		copy(slice, randomSlice)
		slices.Sort(slice)
	}
}

func BenchmarkPDQSortSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, len(sortedSlice))
		copy(slice, sortedSlice)
		slices.Sort(slice)
	}
}
