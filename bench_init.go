package main

import (
	"math/rand"
	"time"
)

const sliceSize = 10000
const bigSliceSize = 100000

var (
	randomSlice []int
	sortedSlice []int

	bigRandomSlice []int
)

func init() {
	randomSlice = generateRandomSlice(sliceSize)
	sortedSlice = generateSortedSlice(sliceSize)

	bigRandomSlice = generateRandomSlice(bigSliceSize)
}

func generateRandomSlice(n int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, n)
	for i := range slice {
		slice[i] = rand.Intn(n)
	}
	return slice
}

func generateSortedSlice(n int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}
	return slice
}
