package main

func checkRecord(n int) (ans int) {
	const mod int = 1e9 + 7
	dp := make([][2][3]int, n+1)
	dp[0][0][0] = 1

	for i := 1; i <= n; i++ {

		for j := 0; j <= 1; j++ {
			for k := 0; k <= 2; k++ {
				dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % mod
			}
		}

		for k := 0; k <= 2; k++ {
			dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % mod
		}

		for j := 0; j <= 1; j++ {
			for k := 1; k <= 2; k++ {
				dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k-1]) % mod
			}
		}
	}
	for j := 0; j <= 1; j++ {
		for k := 0; k <= 2; k++ {
			ans = (ans + dp[n][j][k]) % mod
		}
	}

	return ans
}
