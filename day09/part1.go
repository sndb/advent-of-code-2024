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

	i := 0
	j := len(disk) - 1
	for i < j {
		if disk[i] != space {
			i++
			continue
		}
		if disk[j] == space {
			j--
			continue
		}
		disk[i], disk[j] = disk[j], space
		i++
		j--
	}

	answer := 0
	for i, id := range disk {
		if id == space {
			break
		}
		answer += i * id
	}
	fmt.Println(answer)
}
