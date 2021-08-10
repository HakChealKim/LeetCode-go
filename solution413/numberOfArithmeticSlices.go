package main

func numberOfArithmeticSlices(nums []int) (ans int) {
	n := len(nums)

	if n == 1 {
		return
	}

	delta, t := nums[1] - nums[0], 0

	for i := 2; i < n; i++ {
		if nums[i] - nums[i - 1] == delta {
			t++;
		} else {
			delta, t = nums[i] - nums[i - 1], 0
		}
		ans += t
	}
	return
}
