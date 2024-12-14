package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type robot struct {
	px, py int
	vx, vy int
}

func main() {
	const width = 101
	const height = 103
	const limit = 20000

	robots := []robot{}
	line := bufio.NewScanner(os.Stdin)
	for line.Scan() {
		var px, py, vx, vy int
		fmt.Sscanf(line.Text(), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, robot{px, py, vx, vy})
	}

	answer := 0
	streak := 0
	for second := 1; second < limit; second++ {
		for i, r := range robots {
			robots[i].px = ((r.px+r.vx)%width + width) % width
			robots[i].py = ((r.py+r.vy)%height + height) % height
		}

		sort.Slice(robots, func(i, j int) bool {
			return robots[i].px == robots[j].px && robots[i].py < robots[j].py
		})

		line := 0
		for i := 1; i < len(robots); i++ {
			if robots[i-1].px == robots[i].px && robots[i-1].py == robots[i].py-1 {
				line++
				if line > streak {
					streak = line
					answer = second
				}
			} else {
				line = 0
			}
		}
	}
	fmt.Println(answer)
}
