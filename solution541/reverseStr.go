package main

func reverseStr(s string, k int) string {
	t := []byte(s)
	n := len(t)
	for i := 0; i < n; i += 2 * k {
		j := i
		k1 := min(i+k-1, n-1)
		for j < k1 {
			t[j], t[k1] = t[k1], t[j]
			j++
			k1--
		}
	}
	return string(t)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
