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
