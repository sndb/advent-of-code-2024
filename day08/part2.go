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
				for r1 >= 0 && r1 < n && c1 >= 0 && c1 < n {
					antinodes[[2]int{r1, c1}] = true
					r1 -= dr
					c1 -= dc
				}
				for r2 >= 0 && r2 < n && c2 >= 0 && c2 < n {
					antinodes[[2]int{r2, c2}] = true
					r2 += dr
					c2 += dc
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
