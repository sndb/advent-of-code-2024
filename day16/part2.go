package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"math"
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

	dist1 := dijkstra(grid, sr, sc)
	dist2 := dijkstra(grid, er, ec)
	target := dist1[er][ec]

	answer := 0
	for r := range n {
		for c := range n {
			d := dist1[r][c] + dist2[r][c]
			if d == target || d-1000 == target {
				answer++
			}
		}
	}
	fmt.Println(answer)
}

func dijkstra(grid [][]byte, sr, sc int) [][]int {
	n := len(grid)

	seen := make([][]int, n)
	dist := make([][]int, n)
	for r := range n {
		seen[r] = make([]int, n)
		dist[r] = make([]int, n)
		for c := range n {
			dist[r][c] = math.MaxInt
		}
	}

	queue := &states{{0, sr, sc, 0, 1}}
	for queue.Len() > 0 {
		q := heap.Pop(queue).(state)
		dist[q.r][q.c] = min(dist[q.r][q.c], q.dist)

		d := 1 << (3*q.dr + q.dc + 4)
		if seen[q.r][q.c]&d > 0 {
			continue
		}
		seen[q.r][q.c] |= d

		if nr, nc := q.r+q.dr, q.c+q.dc; grid[nr][nc] != '#' {
			heap.Push(queue, state{q.dist + 1, nr, nc, q.dr, q.dc})
		}
		heap.Push(queue, state{q.dist + 1000, q.r, q.c, -q.dc, q.dr})
		heap.Push(queue, state{q.dist + 1000, q.r, q.c, q.dc, -q.dr})
	}
	return dist
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
