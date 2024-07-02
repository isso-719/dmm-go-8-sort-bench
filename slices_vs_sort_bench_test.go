package main

import (
	"slices"
	"sort"
	"testing"
)

func BenchmarkSlicesPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, len(bigRandomSlice))
		copy(slice, bigRandomSlice)
		slices.Sort(slice)
	}
}

func BenchmarkSortPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, len(bigRandomSlice))
		copy(slice, bigRandomSlice)
		sort.Sort(sort.IntSlice(slice))
	}
}

func tmp() {
	slice := []int{3, 2, 1}
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
}
