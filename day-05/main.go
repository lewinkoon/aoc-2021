package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

// board size
const (
	w = 1000 // width
	h = 1000 // height
)

var board [w * h]int

func coor(a, b string) point {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)

	return point{x, y}
}

func plot(a, b point) {
	dx, dy := b.x-a.x, b.y-a.y

	switch {
	case dx == 0: //vertical
		x := a.x
		ymin, ymax := sort(a.y, b.y)
		for y := ymin; y <= ymax; y++ {
			board[x*w+y]++
		}
	case dy == 0: // horizontal
		y := a.y
		xmin, xmax := sort(a.x, b.x)
		for x := xmin; x <= xmax; x++ {
			board[x*w+y]++
		}
	default: // diagonal
		m := dy / dx
		c := a.y - a.x*m
		xmin, xmax := sort(a.x, b.x)
		for x := xmin; x <= xmax; x++ {
			y := m*x + c
			board[x*w+y]++
		}
	}
}

// avoid negative for loop
func sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// count overlaps between lines
func overlaps() int {
	c := 0
	for i := range board {
		if board[i] > 1 {
			c++
		}
	}
	return c
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := strings.Replace(input.Text(), " -> ", ",", 1)
		args := strings.Split(line, ",")

		a, b := coor(args[0], args[1]), coor(args[2], args[3])
		plot(a, b)
	}

	// print solution
	fmt.Printf("Overlaps: %d\n", overlaps())
}
