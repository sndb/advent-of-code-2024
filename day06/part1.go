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

	result := 0
	dr, dc := -1, 0
	seen := map[[2]int]bool{}
	for {
		state := [2]int{gr, gc}
		if !seen[state] {
			seen[state] = true
			result++
		}

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
	fmt.Println(result)
}
