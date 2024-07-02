package main

func partition(slice []int, n, p, r int) int {
	i := p - 1    // 先頭要素
	x := slice[r] // 基準値(最終idx)
	var tmp int   // 一時変数

	for j := p; j < r; j++ {
		// 比較対象が基準値x以下の場合は入れ替え
		if slice[j] <= x {
			i++
			tmp = slice[i]
			slice[i] = slice[j]
			slice[j] = tmp
		}
	}
	// 基準値を中央に移動
	tmp = slice[i+1]
	slice[i+1] = slice[r]
	slice[r] = tmp

	return i + 1
}

func QuickSort(slice []int, n, p, r int) []int {
	var q int
	if p < r {
		q = partition(slice, n, p, r) // 対象を部分配列に分割
		QuickSort(slice, n, p, q-1)   // 前方部分のクイックソート
		QuickSort(slice, n, q+1, r)   // 後方部分のクイックソート
	}
	return slice
}
