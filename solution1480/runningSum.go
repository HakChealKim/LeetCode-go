package main

func runningSum(nums []int) []int {
	presum := make([]int, len(nums)+1)
	presum[0] = 0

	for i, value := range nums {
		presum[i+1] = presum[i] + value
	}
	return presum[1:]
}

func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
