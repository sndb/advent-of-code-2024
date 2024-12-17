package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	fields := strings.Fields(string(data))

	var program []int
	for _, s := range strings.Split(fields[10], ",") {
		x, _ := strconv.Atoi(s)
		program = append(program, x)
	}

	fmt.Println(search(program, 0, len(program)-1))
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

func search(program []int, acc, pos int) int {
	for i := range 8 {
		val := acc + i
		out := execute(val, 0, 0, program)
		if slices.Equal(out, program[pos:]) {
			if pos == 0 {
				return val
			}
			if a := search(program, val*8, pos-1); a != -1 {
				return a
			}
		}
	}
	return -1
}
