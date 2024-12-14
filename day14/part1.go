package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const width = 101
	const height = 103

	quad := [4]int{}
	line := bufio.NewScanner(os.Stdin)
	for line.Scan() {
		var px, py, vx, vy int
		fmt.Sscanf(line.Text(), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		x := ((px+100*vx)%width + width) % width
		y := ((py+100*vy)%height + height) % height
		cx := width / 2
		cy := height / 2

		if x != cx && y != cy {
			i := 0
			if x < cx {
				i += 2
			}
			if y < cy {
				i += 1
			}
			quad[i]++
		}
	}

	r := 1
	for _, v := range quad {
		r *= v
	}
	fmt.Println(r)
}
