package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../input.txt")

	slices := strings.Split(string(data), "\n")
	prev := -1
	c := 0
	for _, s := range slices {
		n, _ := strconv.Atoi(s)
		if prev != -1 {
			if n > prev {
				c++
			}
		}
		prev = n
	}

	fmt.Printf("Increments: %d", c)
}
