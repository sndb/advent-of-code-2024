package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func nextSecret(x int) int {
	const mask = (1 << 24) - 1
	x = (x ^ x<<6) & mask
	x = (x ^ x>>5) & mask
	x = (x ^ x<<11) & mask
	return x
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	numbers := strings.Fields(string(data))

	answer := 0
	for _, n := range numbers {
		secret, _ := strconv.Atoi(n)
		for range 2000 {
			secret = nextSecret(secret)
		}
		answer += secret
	}
	fmt.Println(answer)
}
