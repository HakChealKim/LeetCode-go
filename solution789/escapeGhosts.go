package main

func escapeGhosts(ghosts [][]int, target []int) bool {
	me := []int{0, 0}
	distance1 := manhattanDistance(me, target)

	for _, ghost := range ghosts {
		distance := manhattanDistance(ghost, target)
		if distance <= distance1 {
			return false
		}
	}
	return true
}

func manhattanDistance(point1, point2 []int) int {
	return abs(point1[0], point2[0]) + abs(point1[1], point2[1])
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
