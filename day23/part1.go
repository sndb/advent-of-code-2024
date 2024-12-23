package main

import (
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strings"
)

func main() {
	type link [2]string

	adj := map[[2]string]bool{}
	nodes := map[string]bool{}

	data, _ := io.ReadAll(os.Stdin)
	links := strings.Fields(string(data))
	for _, l := range links {
		n1, n2 := l[:2], l[3:]
		nodes[n1] = true
		nodes[n2] = true
		adj[link{n1, n2}] = true
		adj[link{n2, n1}] = true
	}

	answer := 0
	list := slices.Collect(maps.Keys(nodes))
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			for k := j + 1; k < len(list); k++ {
				a, b, c := list[i], list[j], list[k]
				if a[0] != 't' && b[0] != 't' && c[0] != 't' {
					continue
				}
				if adj[link{a, b}] && adj[link{a, c}] && adj[link{b, c}] {
					answer++
				}
			}
		}
	}
	fmt.Println(answer)
}
