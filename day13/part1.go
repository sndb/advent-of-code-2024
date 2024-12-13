package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func solve(ax, ay, bx, by, px, py int) int {
	r := math.MaxInt
	for a := range 101 {
		for b := range 101 {
			if ax*a+bx*b == px && ay*a+by*b == py {
				r = min(r, 3*a+b)
			}
		}
	}
	if r == math.MaxInt {
		return 0
	}
	return r
}

func main() {
	answer := 0
	data, _ := io.ReadAll(os.Stdin)
	for _, s := range strings.Split(string(data), "\n\n") {
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(s, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&ax, &ay, &bx, &by, &px, &py)
		answer += solve(ax, ay, bx, by, px, py)
	}
	fmt.Println(answer)
}
