package main

func getIndex(label, row int) int {
	return 1<<(row-1) + 1<<row - 1 - label
}

func pathInZigZagTree(label int) (path []int) {
	row, rowStart := 1, 1
	for rowStart*2 <= label {
		row++
		rowStart *= 2
	}

	if row % 2 == 0 {
		label = getIndex(label, row)
	}

	for row > 0 {
		if row % 2 == 0 {
			path = append(path, getIndex(label, row))
		} else {
			path = append(path, label)
		}
		row--
		label >>= 1
	}

	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}
	return
}
