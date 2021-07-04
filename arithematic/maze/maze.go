package main

import (
	"fmt"
	"os"
)

func read(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	// 定义一个长度为row的[]int slice
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func print(m [][]int) {
	for _, row := range m {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}
}

type point struct {
	x, y int
}

var dir = [4]point{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func (p point) add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

func (p point) isValid(steps [][]int, start point) bool {
	return p.x >= 0 &&
		p.x < len(steps) &&
		p.y >= 0 &&
		p.y < len(steps[0]) &&
		steps[p.x][p.y] == 0 &&
		p != start
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
		for j := range steps[i] {
			if maze[i][j] == 0 {
				steps[i][j] = 0
			} else {
				steps[i][j] = -1
			}
		}
	}

	queue := []point{start}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == end {
			break
		}
		for _, p := range dir {
			newNode := node.add(p)
			if newNode.isValid(steps, start) {
				queue = append(queue, newNode)
				steps[newNode.x][newNode.y] = steps[node.x][node.y] + 1
			}
		}
	}
	return steps
}

func main() {
	maze := read("arithematic/maze/maze.in")
	print(maze)
	steps := walk(maze, point{0, 0}, point{5, 4})
	fmt.Println("result")
	print(steps)
}
