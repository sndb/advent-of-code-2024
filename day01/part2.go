package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ns1 := map[int]int{}
	ns2 := map[int]int{}
	ln := bufio.NewScanner(os.Stdin)
	for ln.Scan() {
		var n1, n2 int
		fmt.Sscanf(ln.Text(), "%d %d", &n1, &n2)
		ns1[n1]++
		ns2[n2]++
	}

	result := 0
	for n, c := range ns1 {
		result += n * c * ns2[n]
	}
	fmt.Println(result)
}
