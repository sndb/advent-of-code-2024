package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	expr := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|don't\(\)`)
	result := 0
	flag := true
	for _, m := range expr.FindAllStringSubmatch(string(data), -1) {
		if s := m[0]; s == "do()" {
			flag = true
		} else if s == "don't()" {
			flag = false
		} else if flag {
			var n1, n2 int
			fmt.Sscanf(s, "mul(%d,%d)", &n1, &n2)
			result += n1 * n2
		}
	}
	fmt.Println(result)
}
