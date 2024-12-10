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

	count := func(r, c int) int {
		seen := make([][]bool, n)
		for r := range n {
			seen[r] = make([]bool, n)
		}

		ret := 0
		queue := [][2]int{{r, c}}
		for len(queue) > 0 {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]

			height := grid[r][c]
			for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= n || nc < 0 || nc >= n || grid[nr][nc] != height+1 {
					continue
				}

				if seen[nr][nc] {
					continue
				}
				seen[nr][nc] = true

				if grid[nr][nc] == '9' {
					ret++
					continue
				}

				queue = append(queue, [2]int{nr, nc})
			}
		}
		return ret
	}

	answer := 0
	for r := range n {
		for c := range n {
			if grid[r][c] == '0' {
				answer += count(r, c)
			}
		}
	}
	fmt.Println(answer)
}
