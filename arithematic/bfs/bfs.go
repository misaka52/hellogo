package main

import "fmt"

type node struct {
	x     int
	y     int
	depth int
}

func bfs(m [][]int) int {
	if m == nil {
		return -1
	}
	row := len(m)
	col := len(m[0])

	trace := make([][]int, row)

	for i := 0; i < row; i++ {
		trace[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if m[i][j] == 0 {
				trace[i][j] = 0
			} else {
				trace[i][j] = -1
			}
		}
	}

	var queue []node
	dir := [4][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	minDepth := -1
	queue = append(queue, node{0, 0, 0})
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		for i := 0; i < len(dir); i++ {
			x := elem.x + dir[i][0]
			y := elem.y + dir[i][1]
			if x >= 0 && x < row && y >= 0 && y < col && trace[x][y] == 0 {
				if x == row-1 && y == col-1 {
					minDepth = elem.depth + 1
					break
				}
				newNode := node{
					x, y, elem.depth + 1,
				}
				queue = append(queue, newNode)
				trace[x][y] = newNode.depth
			}
		}
	}
	fmt.Println("the minDepth is ", minDepth)
	return minDepth
}

func main() {
	bfs([][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	})
	bfs([][]int{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
	})
}
