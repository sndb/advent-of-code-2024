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
		answer += blink(n, 25)
	}
	fmt.Println(answer)
}

func blink(s, n int) int {
	if n == 0 {
		return 1
	}
	if s == 0 {
		return blink(1, n-1)
	}
	if d := digits(s); d%2 == 0 {
		x := 1
		for range d / 2 {
			x *= 10
		}
		return blink(s/x, n-1) + blink(s%x, n-1)
	}
	return blink(s*2024, n-1)
}

func digits(n int) int {
	c := 0
	for n > 0 {
		n /= 10
		c++
	}
	return c
}
