package main

func getMaximumGenerated(n int) (ans int) {
	if n == 0 {
		return
	}
	if n == 1 {
		ans = 1
		return
	}

	a := make([]int, n+1)
	a[0] = 0
	a[1] = 1

	for i := 1; i <= n/2; i++ {
		a[2*i] = a[i]
		if 2*i+1 < n {
			a[2*i+1] = a[i] + a[i+1]
			t := max(a[2*i], a[2*i+1])
			ans = max(t, ans)
		} else {
			ans = max(a[2*i], ans)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
