package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	split := strings.Split(string(data), "\n\n")

	patterns := map[string]bool{}
	for _, p := range strings.Split(split[0], ", ") {
		patterns[p] = true
	}

	answer := 0
	for _, d := range strings.Fields(split[1]) {
		answer += possible(patterns, d)
	}
	fmt.Println(answer)
}

var cache = map[string]int{"": 1}

func possible(patterns map[string]bool, design string) int {
	if p, ok := cache[design]; ok {
		return p
	}

	p := 0
	for n := 1; n <= len(design); n++ {
		if patterns[design[:n]] {
			p += possible(patterns, design[n:])
		}
	}
	cache[design] = p
	return p
}
