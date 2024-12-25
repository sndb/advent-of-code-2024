package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	parts := strings.Split(string(data), "\n\n")
	part1 := strings.Fields(parts[0])
	part2 := strings.Fields(parts[1])

	state := map[string]int{}
	for i := 0; i < len(part1); i += 2 {
		wire := part1[i][:3]
		bit := part1[i+1][0] - '0'
		state[wire] = int(bit)
	}

	type expr struct {
		x, y string
		op   string
	}

	graph := map[string]expr{}
	for i := 0; i < len(part2); i += 5 {
		x := part2[i]
		f := part2[i+1]
		y := part2[i+2]
		z := part2[i+4]
		graph[z] = expr{x, y, f}
	}

	funcs := map[string]func(int, int) int{
		"AND": func(x, y int) int { return x & y },
		"OR":  func(x, y int) int { return x | y },
		"XOR": func(x, y int) int { return x ^ y },
	}

	var eval func(string) int
	eval = func(z string) int {
		if v, ok := state[z]; ok {
			return v
		}
		expr := graph[z]
		state[z] = funcs[expr.op](eval(expr.x), eval(expr.y))
		return state[z]
	}

	i := 0
	v := 0
	for {
		z := fmt.Sprintf("z%02d", i)
		if _, ok := graph[z]; !ok {
			break
		}
		v |= eval(z) << i
		i++
	}
	fmt.Println(v)
}
