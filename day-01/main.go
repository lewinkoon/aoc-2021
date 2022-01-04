package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	slices := strings.Split(string(data), "\n")

	fmt.Printf("Part one: %d\n", p1(slices))
	fmt.Printf("Part two: %d\n", p2(slices))
}

func p1(inp []string) int {

	prev := -1
	inc := 0
	for _, s := range inp {
		n, _ := strconv.Atoi(s)
		if prev != -1 {
			if n > prev {
				inc++
			}
		}
		prev = n
	}

	return inc
}

func p2(inp []string) int {
	prev := -1
	inc := 0

	for i := range inp {

		// terminate loop
		if i == len(inp)-2 {
			break
		}

		n, _ := strconv.Atoi(inp[i])
		n1, _ := strconv.Atoi(inp[i+1])
		n2, _ := strconv.Atoi(inp[i+2])
		tot := n + n1 + n2

		if prev != -1 {
			if tot > prev {
				inc++
			}
		}

		// update previous value
		prev = tot
	}

	return inc
}
