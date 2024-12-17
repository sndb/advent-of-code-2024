package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	fields := strings.Fields(string(data))

	a, _ := strconv.Atoi(fields[2])
	b, _ := strconv.Atoi(fields[5])
	c, _ := strconv.Atoi(fields[8])

	var program []int
	for _, s := range strings.Split(fields[10], ",") {
		x, _ := strconv.Atoi(s)
		program = append(program, x)
	}

	out := execute(a, b, c, program)
	for i, v := range out {
		fmt.Print(v)
		if i == len(out)-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(",")
		}
	}
}

func execute(a, b, c int, program []int) []int {
	var out []int
	for ptr := 0; ptr < len(program); ptr += 2 {
		in := program[ptr]
		op := program[ptr+1]

		var combo int
		switch op {
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		default:
			combo = op
		}

		switch in {
		case 0: // adv
			a >>= combo
		case 1: // bxl
			b ^= op
		case 2: // bst
			b = combo & 7
		case 3: // jnz
			if a != 0 {
				ptr = op - 2
			}
		case 4: // bxc
			b ^= c
		case 5: // out
			out = append(out, combo%8)
		case 6: // bdv
			b = a >> combo
		case 7: // cdv
			c = a >> combo
		}
	}
	return out
}
