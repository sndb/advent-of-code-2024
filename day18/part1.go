package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	size  = 71
	steps = 1024
)

type pos struct {
	r, c int
}

func main() {
	bytes := map[pos]bool{}
	data, _ := io.ReadAll(os.Stdin)
	for _, line := range strings.Fields(string(data))[:steps] {
		var p pos
		fmt.Sscanf(line, "%d,%d", &p.r, &p.c)
		bytes[p] = true
	}

	answer := 0
	seen := map[pos]bool{}
	queue := []pos{{0, 0}}
	for {
		front := []pos{}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if p.r == size-1 && p.c == size-1 {
				fmt.Println(answer)
				return
			}

			for _, q := range []pos{
				{p.r - 1, p.c},
				{p.r + 1, p.c},
				{p.r, p.c - 1},
				{p.r, p.c + 1},
			} {
				if q.r >= 0 && q.r < size && q.c >= 0 && q.c < size && !bytes[q] && !seen[q] {
					seen[q] = true
					front = append(front, q)
				}
			}
		}
		queue = front
		answer++
	}
}
