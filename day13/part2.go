package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(ax, ay, bx, by, px, py int) int {
	d := ax*by - ay*bx
	a := (px*by - py*bx) / d
	b := (py*ax - px*ay) / d
	if px != a*ax+b*bx || py != a*ay+b*by {
		return 0
	}
	return 3*a + b
}

func main() {
	answer := 0
	data, _ := io.ReadAll(os.Stdin)
	for _, s := range strings.Split(string(data), "\n\n") {
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(s, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&ax, &ay, &bx, &by, &px, &py)
		answer += solve(ax, ay, bx, by, 10000000000000+px, 10000000000000+py)
	}
	fmt.Println(answer)
}
