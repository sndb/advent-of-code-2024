package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var ns1, ns2 []int
	ln := bufio.NewScanner(os.Stdin)
	for ln.Scan() {
		var n1, n2 int
		fmt.Sscanf(ln.Text(), "%d %d", &n1, &n2)
		ns1 = append(ns1, n1)
		ns2 = append(ns2, n2)
	}
	sort.Ints(ns1)
	sort.Ints(ns2)

	result := 0
	for i, n1 := range ns1 {
		n2 := ns2[i]
		diff := n2 - n1
		if diff < 0 {
			diff = -diff
		}
		result += diff
	}
	fmt.Println(result)
}
