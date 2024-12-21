package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row int
	col int
}

type path struct {
	src byte
	dst byte
}

func computeSeqs(keypad [][]byte) map[path][]string {
	coords := map[byte]coord{}
	rows := len(keypad)
	cols := len(keypad[0])
	for r := range rows {
		for c := range cols {
			if keypad[r][c] != ' ' {
				coords[keypad[r][c]] = coord{r, c}
			}
		}
	}

	result := map[path][]string{}
	for k1, c1 := range coords {
		for k2, c2 := range coords {
			if k1 == k2 {
				result[path{k1, k2}] = []string{"A"}
				continue
			}

			seqs := []string{}
			optimal := math.MaxInt

			type entry struct {
				coord coord
				path  string
			}
			queue := []entry{{c1, ""}}
			for len(queue) > 0 {
				e := queue[0]
				queue = queue[1:]

				if e.coord == c2 {
					if optimal < len(e.path) {
						continue
					}
					optimal = len(e.path)
					seqs = append(seqs, e.path+"A")
					continue
				}

				for _, d := range []string{"^", "v", "<", ">"} {
					var dr, dc int
					switch d {
					case "^":
						dr, dc = -1, 0
					case "v":
						dr, dc = 1, 0
					case "<":
						dr, dc = 0, -1
					case ">":
						dr, dc = 0, 1
					}

					nr := e.coord.row + dr
					nc := e.coord.col + dc
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}
					if keypad[nr][nc] == ' ' || len(e.path) > optimal {
						continue
					}
					queue = append(queue, entry{coord{nr, nc}, e.path + d})
				}
			}
			result[path{k1, k2}] = seqs
		}
	}
	return result
}

var numKeypad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var numSeqs = computeSeqs(numKeypad)

var dirKeypad = [][]byte{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
}

var dirSeqs = computeSeqs(dirKeypad)

func initialSequences(code string, seqs map[path][]string) []string {
	result := seqs[path{'A', code[0]}]
	for i := 1; i < len(code); i++ {
		next := []string{}
		for _, p := range result {
			for _, s := range seqs[path{code[i-1], code[i]}] {
				next = append(next, p+s)
			}
		}
		result = next
	}
	return result
}

var cache = map[cacheKey]int{}

type cacheKey struct {
	seq   string
	depth int
}

func computeLength(seq string, depth int) int {
	k := cacheKey{seq, depth}
	if v, ok := cache[k]; ok {
		return v
	}
	cache[k] = computeLength1(seq, depth)
	return cache[k]
}

func computeLength1(seq string, depth int) int {
	if depth == 1 {
		sum := len(dirSeqs[path{'A', seq[0]}][0])
		for i := 1; i < len(seq); i++ {
			sum += len(dirSeqs[path{seq[i-1], seq[i]}][0])
		}
		return sum
	}

	length := math.MaxInt
	for _, subseq := range dirSeqs[path{'A', seq[0]}] {
		length = min(length, computeLength(subseq, depth-1))
	}
	for i := 1; i < len(seq); i++ {
		l := math.MaxInt
		for _, subseq := range dirSeqs[path{seq[i-1], seq[i]}] {
			l = min(l, computeLength(subseq, depth-1))
		}
		length += l
	}
	return length
}

func main() {
	answer := 0
	data, _ := io.ReadAll(os.Stdin)
	for _, code := range strings.Fields(string(data)) {
		n, _ := strconv.Atoi(code[:3])
		l := math.MaxInt
		for _, s := range initialSequences(code, numSeqs) {
			l = min(l, computeLength(s, 25))
		}
		answer += n * l
	}
	fmt.Println(answer)
}
