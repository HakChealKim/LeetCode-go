package main

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	dp := make([][]int, k+2)

	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0

	for i := 1; i <= k+1; i++ {
		for _, flight := range flights {
			from := flight[0]
			to := flight[1]
			cost := flight[2]
			dp[i][to] = min(dp[i-1][from]+cost, dp[i][to])
		}
	}

	ans := inf

	for i := 0; i <= k+1; i++ {
		ans = min(dp[i][dst], ans)
	}

	if ans == inf {
		return -1
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
