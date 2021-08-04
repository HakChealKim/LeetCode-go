package main

import "sort"

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	sortedNums := append([]int(nil), nums...)
	sort.Ints(sortedNums)
	left, right := 0, len(nums) - 1
	for nums[left] == sortedNums[left] {
		left++
	}

	for nums[right] == sortedNums[right] {
		right--
	}

	return right - left + 1
}
