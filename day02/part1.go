package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		if safe(report) {
			result++
		}
	}
	fmt.Println(result)
}
