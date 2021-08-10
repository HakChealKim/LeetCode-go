package main

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a) - 1]
	h.IntSlice = a[:len(a) - 1]
	return v
}

func nthSuperUglyNumber(n int, primes []int) (uglyNum int) {
	seen := map[int]bool{1:true}
	h := &hp{[]int{1}}
	for i := 0; i < n; i++ {
		uglyNum = heap.Pop(h).(int)
		for _, prime := range primes {
			if next := uglyNum * prime; !seen[next] {
				seen[next] = true
				heap.Push(h, next)
			}
		}
	}
	return
}
