package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	west = iota
	east
	north
	south

	dirs
)

var dr = [dirs]int{north: -1, south: 1}
var dc = [dirs]int{west: -1, east: 1}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	grid := bytes.Fields(data)
	n := len(grid)

	seen := make([][]bool, n)
	for r := range n {
		seen[r] = make([]bool, n)
	}

	count := func(r, c int) (area, sides int) {
		plant := grid[r][c]

		fence := make([][][4]bool, n)
		for r := range n {
			fence[r] = make([][4]bool, n)
		}

		queue := [][2]int{{r, c}}
		for len(queue) > 0 {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			area++

			for d := range dirs {
				nr, nc := r+dr[d], c+dc[d]
				if nr < 0 || nr >= n || nc < 0 || nc >= n || grid[nr][nc] != plant {
					fence[r][c][d] = true
				} else if !seen[nr][nc] {
					seen[nr][nc] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}

		for r := range n {
			for c := range n {
				if fence[r][c][north] && (c == 0 || !fence[r][c-1][north]) {
					sides++
				}
				if fence[r][c][south] && (c == 0 || !fence[r][c-1][south]) {
					sides++
				}
				if fence[r][c][west] && (r == 0 || !fence[r-1][c][west]) {
					sides++
				}
				if fence[r][c][east] && (r == 0 || !fence[r-1][c][east]) {
					sides++
				}
			}
		}
		return
	}

	answer := 0
	for r := range n {
		for c := range n {
			if !seen[r][c] {
				seen[r][c] = true
				a, p := count(r, c)
				answer += a * p
			}
		}
	}
	fmt.Println(answer)
}
