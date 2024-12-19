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
		if possible(patterns, d) {
			answer++
		}
	}
	fmt.Println(answer)
}

func possible(patterns map[string]bool, design string) bool {
	if design == "" {
		return true
	}
	for n := 1; n <= len(design); n++ {
		if patterns[design[:n]] && possible(patterns, design[n:]) {
			return true
		}
	}
	return false
}
