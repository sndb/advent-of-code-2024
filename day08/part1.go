package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	grid := strings.Fields(string(data))
	n := len(grid)

	antennas := map[byte][][2]int{}
	for r := range n {
		for c := range n {
			if a := grid[r][c]; a != '.' {
				antennas[a] = append(antennas[a], [2]int{r, c})
			}
		}
	}

	antinodes := map[[2]int]bool{}
	for _, rcs := range antennas {
		for i, rc1 := range rcs {
			for _, rc2 := range rcs[i+1:] {
				r1, c1 := rc1[0], rc1[1]
				r2, c2 := rc2[0], rc2[1]
				dr, dc := r2-r1, c2-c1
				antinodes[[2]int{r1 - dr, c1 - dc}] = true
				antinodes[[2]int{r2 + dr, c2 + dc}] = true
			}
		}
	}

	answer := 0
	for rc := range antinodes {
		r, c := rc[0], rc[1]
		if r >= 0 && r < n && c >= 0 && c < n {
			answer++
		}
	}
	fmt.Println(answer)
}
