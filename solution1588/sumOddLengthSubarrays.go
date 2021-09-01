package main

func sumOddLengthSubarrays(arr []int) (ans int) {
	n := len(arr)
	presum := make([]int, n+1)
	presum[0] = arr[0]

	for i := 1; i <= n; i++ {
		presum[i] = arr[i-1] + presum[i-1]
	}

	for i := 1; i <= n; i += 2 {
		for j := 1; j < n; j++ {
			if j-i >= 0 {
				ans += presum[j] - presum[j-i]
				println(ans)
			}
		}
	}

	return

}
