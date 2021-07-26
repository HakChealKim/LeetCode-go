package main

import "sort"

func minOperations(target []int, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)

	for i, value := range target {
		pos[value] = i
	}

	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}

	return n - len(d)
}
