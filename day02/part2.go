package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func permutations(report []int) [][]int {
	r := [][]int{report}
	for i := 0; i < len(report); i++ {
		p := make([]int, len(report)-1)
		copy(p[:i], report[:i])
		copy(p[i:], report[i+1:])
		r = append(r, p)
	}
	return r
}

func safe(report []int) bool {
	sign := 1
	if report[1] < report[0] {
		sign = -1
	}
	for i := 0; i < len(report)-1; i++ {
		diff := (report[i+1] - report[i]) * sign
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func main() {
	result := 0
	ln := bufio.NewScanner(os.Stdin)
	for ln.Scan() {
		var report []int
		for _, s := range strings.Fields(ln.Text()) {
			n, _ := strconv.Atoi(s)
			report = append(report, n)
		}
		for _, p := range permutations(report) {
			if safe(p) {
				result++
				break
			}
		}
	}
	fmt.Println(result)
}
