package main

type matrix [3][3]int

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	m := matrix{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}

	res := m.pow(n)
	return res[0][2]
}

func (a matrix) multiply(b matrix) matrix {
	c := matrix{}
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] += v * b[k][j];
			}
		}
	}
	return c
}

func (a matrix) pow(n int) matrix {
	res := matrix{}
	for i := range res {
		res[i][i] = 1
	}

	for ; n > 0; n >>= 1 {
		if n & 1 > 0{
			res = res.multiply(a)
		}
		a = a.multiply(a)
	}
	return res
}
