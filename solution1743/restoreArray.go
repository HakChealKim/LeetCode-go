package main

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	pairs := make(map[int][]int, n)

	for _, adjacent := range adjacentPairs {
		v, w := adjacent[0], adjacent[1]
		pairs[v] = append(pairs[v], w)
		pairs[w] = append(pairs[w], v)
	}

	ans := make([]int, n)
	for val, adj := range  pairs {
		if len(adj) == 1 {
			ans[0] = val
			break
		}
	}

	ans[1] = pairs[ans[0]][0]
	for i := 2;i < n; i++  {
		adj := pairs[ans[i-1]]
		if ans[i-2] == adj[0] {
			ans[i] = adj[1]
		} else {
			ans[i] = adj[0]
		}
	}
	return ans
}
