package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"os"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	grid := bytes.Fields(data)
	n := len(grid)

	var sr, sc, er, ec int
	for r := range n {
		for c := range n {
			switch grid[r][c] {
			case 'S':
				sr, sc = r, c
			case 'E':
				er, ec = r, c
			}
		}
	}

	seen := make([][]int, n)
	for r := range n {
		seen[r] = make([]int, n)
	}

	queue := &states{{0, sr, sc, 0, 1}}
	for {
		s := heap.Pop(queue).(state)
		if s.r == er && s.c == ec {
			fmt.Println(s.dist)
			return
		}

		d := 1 << (3*s.dr + s.dc + 4)
		if seen[s.r][s.c]&d > 0 {
			continue
		}
		seen[s.r][s.c] |= d

		if nr, nc := s.r+s.dr, s.c+s.dc; grid[nr][nc] != '#' {
			heap.Push(queue, state{s.dist + 1, nr, nc, s.dr, s.dc})
		}
		heap.Push(queue, state{s.dist + 1000, s.r, s.c, -s.dc, s.dr})
		heap.Push(queue, state{s.dist + 1000, s.r, s.c, s.dc, -s.dr})
	}
}

type state struct {
	dist   int
	r, c   int
	dr, dc int
}

type states []state

func (h states) Len() int           { return len(h) }
func (h states) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h states) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *states) Push(x any) {
	*h = append(*h, x.(state))
}

func (h *states) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
