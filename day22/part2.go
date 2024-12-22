package main

import (
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
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

func seqPrices(secret int) map[[4]int]int {
	seq := [4]int{}
	prices := map[[4]int]int{}
	for i := range 2000 {
		next := nextSecret(secret)
		p1, p2 := secret%10, next%10
		seq = [4]int{seq[1], seq[2], seq[3], p2 - p1}
		secret = next

		if _, ok := prices[seq]; !ok && i >= 3 {
			prices[seq] = p2
		}
	}
	return prices
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	numbers := strings.Fields(string(data))

	total := map[[4]int]int{}
	for _, s := range numbers {
		n, _ := strconv.Atoi(s)
		for seq, price := range seqPrices(n) {
			total[seq] += price
		}
	}
	fmt.Println(slices.Max(slices.Collect(maps.Values(total))))
}
