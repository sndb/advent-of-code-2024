package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)

	answer := 0
	for _, s := range strings.Fields(string(data)) {
		n, _ := strconv.Atoi(s)
		answer += blink(n, 75)
	}
	fmt.Println(answer)
}

var cache = map[[2]int]int{}

func blink(s, n int) int {
	k := [2]int{s, n}
	v, ok := cache[k]
	if ok {
		return v
	}

	if n == 0 {
		v = 1
	} else if s == 0 {
		v = blink(1, n-1)
	} else if d := digits(s); d%2 == 0 {
		x := 1
		for range d / 2 {
			x *= 10
		}
		v = blink(s/x, n-1) + blink(s%x, n-1)
	} else {
		v = blink(s*2024, n-1)
	}

	cache[k] = v
	return v
}

func digits(n int) int {
	c := 0
	for n > 0 {
		n /= 10
		c++
	}
	return c
}
