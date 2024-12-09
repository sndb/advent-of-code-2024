package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const space = -1

func main() {
	data, _ := io.ReadAll(os.Stdin)

	disk := []int{}
	for i, c := range bytes.TrimSpace(data) {
		n := c - '0'
		if i%2 == 0 {
			// file
			for range n {
				disk = append(disk, i/2)
			}
		} else {
			// space
			for range n {
				disk = append(disk, space)
			}
		}
	}

	end := len(disk) - 1
	for end > 0 {
		if disk[end] == space {
			end--
			continue
		}

		start, id := end, disk[end]
		for start > 0 && disk[start] == id {
			start--
		}
		start++

		i := 0
		for i < start {
			if disk[i] != space {
				i++
				continue
			}

			j := i
			for disk[j] == space {
				j++
			}

			diff := j - i
			size := end - start + 1
			if diff < size {
				i = j + 1
				continue
			}

			for k := 0; k < size; k++ {
				disk[i+k] = id
				disk[start+k] = space
			}
			break
		}
		end = start - 1
	}

	answer := 0
	for i, id := range disk {
		if id != space {
			answer += i * id
		}
	}
	fmt.Println(answer)
}
