package main

import (
	"container/heap"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	h := hp{}
	for i, row := range mat {
		score := sort.Search(len(row), func(j int) bool {
			return row[j] == 0
		})
		h = append(h, pair{score, i})
	}

	heap.Init(&h)
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).idx
	}
	return ans
}

type pair struct{ score, idx int }
type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.score < b.score || a.score == b.score && a.idx < b.idx
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a) - 1]
	*h = a[:len(a) - 1]
	return v
}
