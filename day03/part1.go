package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	expr := regexp.MustCompile(`mul\(\d+,\d+\)`)
	result := 0
	for _, m := range expr.FindAllStringSubmatch(string(data), -1) {
		var n1, n2 int
		fmt.Sscanf(m[0], "mul(%d,%d)", &n1, &n2)
		result += n1 * n2
	}
	fmt.Println(result)
}
