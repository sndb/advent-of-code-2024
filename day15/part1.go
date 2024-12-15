package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	split := bytes.Split(data, []byte("\n\n"))
	grid := bytes.Fields(split[0])
	moves := split[1]
	n := len(grid)

	var rr, rc int
outer:
	for r := range n {
		for c := range n {
			if grid[r][c] == '@' {
				grid[r][c] = '.'
				rr, rc = r, c
				break outer
			}
		}
	}

	for _, move := range moves {
		dr, dc := 0, 0
		switch move {
		case '^':
			dr = -1
		case 'v':
			dr = 1
		case '<':
			dc = -1
		case '>':
			dc = 1
		default:
			continue
		}

		stack := []byte{}
		tr, tc := rr+dr, rc+dc
		for grid[tr][tc] != '#' {
			x := grid[tr][tc]
			if x == '.' {
				rr, rc = rr+dr, rc+dc
				for len(stack) > 0 {
					l := len(stack)
					grid[tr][tc] = stack[l-1]
					stack = stack[:l-1]
					tr -= dr
					tc -= dc
				}
				grid[tr][tc] = '.'
				break
			}
			stack = append(stack, x)
			tr += dr
			tc += dc
		}
	}

	answer := 0
	for r := range n {
		for c := range n {
			if grid[r][c] == 'O' {
				answer += r*100 + c
			}
		}
	}
	fmt.Println(answer)
}
