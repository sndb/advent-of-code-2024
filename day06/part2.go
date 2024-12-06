package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	grid := bytes.Fields(input)
	n := len(grid)

	inside := func(r, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < n
	}

	var gr, gc int
	for r := range n {
		if c := bytes.IndexByte(grid[r], '^'); c != -1 {
			gr, gc = r, c
		}
	}

	loops := func() bool {
		gr, gc := gr, gc
		dr, dc := -1, 0
		seen := map[[4]int]bool{}
		for {
			state := [4]int{gr, gc, dr, dc}
			if seen[state] {
				return true
			}
			seen[state] = true

			nr, nc := gr+dr, gc+dc
			if !inside(nr, nc) {
				break
			}
			if grid[nr][nc] == '#' {
				dr, dc = dc, -dr
			} else {
				gr += dr
				gc += dc
			}
		}
		return false
	}

	result := 0
	for r := range n {
		for c := range n {
			if grid[r][c] != '.' {
				continue
			}

			grid[r][c] = '#'
			if loops() {
				result++
			}
			grid[r][c] = '.'
		}
	}
	fmt.Println(result)
}
