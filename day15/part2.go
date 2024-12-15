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
	moves := split[1]

	var rr, rc int
	var grid [][]byte
	for r, s := range bytes.Fields(split[0]) {
		var row []byte
		for c, x := range s {
			if x == '@' {
				rr, rc = r, c*2
				row = append(row, '.', '.')
			} else if x == 'O' {
				row = append(row, '[', ']')
			} else {
				row = append(row, x, x)
			}
		}
		grid = append(grid, row)
	}
	rows := len(grid)
	cols := len(grid[0])

	push := func(or, oc, dr, dc int) map[[2]int]byte {
		seen := map[[2]int]bool{}
		queue := [][2]int{{or, oc}}
		for len(queue) > 0 {
			q := queue[0]
			queue = queue[1:]

			if seen[q] {
				continue
			}
			seen[q] = true

			r, c := q[0], q[1]
			switch grid[r][c] {
			case '#':
				return nil
			case '[':
				queue = append(queue, [2]int{r, c + 1}, [2]int{r + dr, c + dc})
			case ']':
				queue = append(queue, [2]int{r, c - 1}, [2]int{r + dr, c + dc})
			}
		}
		ret := map[[2]int]byte{}
		for k := range seen {
			if x := grid[k[0]][k[1]]; x != '.' {
				ret[k] = x
			}
		}
		return ret
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

		if boxes := push(rr+dr, rc+dc, dr, dc); boxes != nil {
			for k := range boxes {
				grid[k[0]][k[1]] = '.'
			}
			for k, v := range boxes {
				grid[k[0]+dr][k[1]+dc] = v
			}
			rr += dr
			rc += dc
		}
	}

	answer := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '[' {
				answer += r*100 + c
			}
		}
	}
	fmt.Println(answer)
}
