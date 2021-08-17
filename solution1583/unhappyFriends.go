package main

func unhappyFriends(n int, preferences [][]int, pairs [][]int) (ans int) {
	rank := make([][]int, n)
	for i := 0; i < n; i++ {
		rank[i] = make([]int, n)
		for j := 0; j < n-1; j++ {
			rank[i][preferences[i][j]] = j
		}
	}

	couple := make([]int, n)
	for _, val := range pairs {
		couple[val[0]] = val[1]
		couple[val[1]] = val[0]
	}

	for p1, p2 := range couple {
		index := rank[p1][p2]
		for _, u := range preferences[p1][:index] {
			v := couple[u]
			if rank[u][p1] < rank[u][v] {
				ans++
				break
			}
		}
	}
	return
}
