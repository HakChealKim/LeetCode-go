package main

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	queue := []tuple{}
	seen := make([][]bool, n)

	for i := range seen {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		queue = append(queue, tuple{
			u:    i,
			mask: 1 << i,
			dist: 0,
		})
	}

	for {
		v := queue[0]
		queue = queue[1:]

		if v.mask == 1<<n-1 {
			return v.dist
		}

		for _, val := range graph[v.u] {
			maskV := v.mask | 1<<val
			if !seen[val][maskV] {
				queue = append(queue, tuple{
					u:    val,
					mask: maskV,
					dist: v.dist + 1,
				})

				seen[val][maskV] = true
			}
		}
	}

}
