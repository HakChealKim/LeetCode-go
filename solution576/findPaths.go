package main

const mod int = 1e9 + 7

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) (ans int) {
	dp := make([][][]int, maxMove + 1)

	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	dp[0][startRow][startColumn] = 1
	for i := 0; i < maxMove; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				count := dp[i][j][k]
				if count > 0 {
					for _, dir := range dirs {
						x := j + dir.x
						y := k + dir.y
						if inArea(x, y, m, n) {
							dp[i+1][x][y] = (dp[i+1][x][y] + count) % mod
						} else {
							ans = (ans + count) % mod
						}
					}
				}
			}
		}
	}
	return
}

func inArea(x, y, m, n int) bool {
	if x >= 0 && x < m && y >= 0 && y < n {
		return true
	}
	return false
}
