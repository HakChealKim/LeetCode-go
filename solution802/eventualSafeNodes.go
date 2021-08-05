package main

func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	color := make([]int, n)
	var isSafe func(int) bool

	isSafe = func(x int) bool {
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1
		for _, val := range graph[x] {
			if !isSafe(val) {
				return false
			}
		}
		color[x] = 2
		return true
	}

	for i := 0; i < n; i++ {
		if isSafe(i) {
			ans = append(ans, i)
		}
	}
	return
}