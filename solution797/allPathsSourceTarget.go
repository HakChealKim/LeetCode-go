package main

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	stack := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int(nil), stack...))
			return
		}
		for _, y := range graph[x] {
			stack = append(stack, y)
			dfs(y)
			stack = stack[:len(stack)-1]
		}
	}

	dfs(0)
	return
}
