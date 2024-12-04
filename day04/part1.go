package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	const word = "XMAS"

	data, _ := io.ReadAll(os.Stdin)
	grid := bytes.Fields(data)
	rows := len(grid)
	cols := len(grid[0])

	search := func(or, oc int) int {
		ret := 0
		drs := []int{-1, -1, -1, 0, 0, 1, 1, 1}
		dcs := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	outer:
		for di, dr := range drs {
			dc := dcs[di]
			for i := range word {
				r := or + dr*i
				c := oc + dc*i
				if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != word[i] {
					continue outer
				}
			}
			ret++
		}
		return ret
	}

	result := 0
	for r := range rows {
		for c := range cols {
			result += search(r, c)
		}
	}
	fmt.Println(result)
}
