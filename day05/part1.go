package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	parts := strings.Split(string(input), "\n\n")
	rules, updates := parts[0], parts[1]

	var lt [100][100]bool
	for _, rule := range strings.Fields(rules) {
		var p1, p2 int
		fmt.Sscanf(rule, "%d|%d", &p1, &p2)
		lt[p1][p2] = true
	}

	result := 0
updates:
	for _, update := range strings.Fields(updates) {
		var pages []int
		for _, s := range strings.Split(update, ",") {
			p, _ := strconv.Atoi(s)
			pages = append(pages, p)
		}
		for i, p1 := range pages {
			for _, p2 := range pages[i+1:] {
				if lt[p2][p1] {
					continue updates
				}
			}
		}
		result += pages[len(pages)/2]
	}
	fmt.Println(result)
}
