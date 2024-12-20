package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type pos struct {
	r int
	c int
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	grid := bytes.Fields(data)
	n := len(grid)

	var er, ec int
	dist := make([][]int, n)
	for r := range n {
		dist[r] = make([]int, n)
		for c := range n {
			if grid[r][c] == 'E' {
				er, ec = r, c
			} else {
				dist[r][c] = -1
			}
		}
	}

	steps := 1
	queue := []pos{{er, ec}}
	for len(queue) > 0 {
		frontier := []pos{}
		for len(queue) > 0 {
			q := queue[0]
			r, c := q.r, q.c
			queue = queue[1:]

			for _, p := range []pos{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
				if dist[p.r][p.c] == -1 && grid[p.r][p.c] != '#' {
					dist[p.r][p.c] = steps
					frontier = append(frontier, p)
				}
			}
		}
		queue = frontier
		steps++
	}

	answer := 0
	for r1 := range n {
		for c1 := range n {
			for _, p := range []pos{
				{r1 + 2, c1}, {r1 + 1, c1 + 1}, {r1 + 1, c1 - 1},
				{r1 - 2, c1}, {r1 - 1, c1 + 1}, {r1 - 1, c1 - 1},
				{r1, c1 + 2}, {r1, c1 - 2},
			} {
				r2, c2 := p.r, p.c
				if r2 < 0 || r2 >= n || c2 < 0 || c2 >= n {
					continue
				}
				if dist[r1][c1] == -1 || dist[r2][c2] == -1 {
					continue
				}
				if dist[r1][c1]-dist[r2][c2]-2 >= 100 {
					answer++
				}
			}
		}
	}
	fmt.Println(answer)
}
