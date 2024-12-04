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
	rows := len(grid)
	cols := len(grid[0])

	match := func(r, c int) bool {
		g00, g01, g10, g11 := grid[r][c], grid[r][c+2], grid[r+2][c], grid[r+2][c+2]
		return grid[r+1][c+1] == 'A' &&
			(g00 == 'M' && g11 == 'S' || g00 == 'S' && g11 == 'M') &&
			(g01 == 'M' && g10 == 'S' || g01 == 'S' && g10 == 'M')
	}

	result := 0
	for r := range rows - 2 {
		for c := range cols - 2 {
			if match(r, c) {
				result++
			}
		}
	}
	fmt.Println(result)
}
