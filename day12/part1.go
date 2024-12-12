package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	grid := bytes.Fields(data)
	n := len(grid)

	seen := make([][]bool, n)
	for r := range n {
		seen[r] = make([]bool, n)
	}

	count := func(r, c int) (area, perimeter int) {
		plant := grid[r][c]
		queue := [][2]int{{r, c}}
		for len(queue) > 0 {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			area++

			for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= n || nc < 0 || nc >= n || grid[nr][nc] != plant {
					perimeter++
				} else if !seen[nr][nc] {
					seen[nr][nc] = true
					queue = append(queue, [2]int{nr, nc})
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
