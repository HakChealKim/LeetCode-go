package main

func countArrangement(n int) (ans int) {
	visited := make([]bool, n+1)
	match := make([][]int, n+1)

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}

	var backtrack func(int)
	backtrack = func(index int) {
		if index == n + 1 {
			ans++
			return
		}
		for _, x := range match[index] {
			if !visited[x] {
				visited[x] = true
				backtrack(index + 1)
				visited[x] = false
			}
		}
	}

	backtrack(1)
	return
}
