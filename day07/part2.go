package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := 0
	lines := bufio.NewScanner(os.Stdin)
	for lines.Scan() {
		fields := strings.Fields(lines.Text())
		result, _ := strconv.Atoi(fields[0][:len(fields[0])-1])
		var nums []int
		for _, s := range fields[1:] {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
		if possible(result, nums) {
			answer += result
		}
	}
	fmt.Println(answer)
}

func possible(result int, nums []int) bool {
	vals := map[int]bool{nums[0]: true}
	nums = nums[1:]
	for len(nums) > 0 {
		n := nums[0]
		next := map[int]bool{}
		for v := range vals {
			next[v+n] = true
			next[v*n] = true
			next[concat(v, n)] = true
		}
		nums = nums[1:]
		vals = next
	}
	return vals[result]
}

func concat(n1, n2 int) int {
	for m := n2; m > 0; m /= 10 {
		n1 *= 10
	}
	return n1 + n2
}
