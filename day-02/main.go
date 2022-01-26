package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("test.txt")

	lines := strings.Split(string(data), "\n")

	fmt.Printf("Part one: %d\n", p1(lines))
	fmt.Printf("Part two: %d\n", p2(lines))
}

func p1(input []string) int {

	x, y := 0, 0
	for _, line := range input {
		cmd := strings.Split(line, " ")
		ins := cmd[0]
		num, _ := strconv.Atoi(cmd[1])

		switch ins {
		case "forward":
			x += num
		case "up":
			y -= num
		case "down":
			y += num
		default:
			fmt.Println("Error")
		}
	}

	return x * y
}

func p2(input []string) int {

	x, y, aim := 0, 0, 0
	for _, line := range input {
		cmd := strings.Split(line, " ")
		ins := cmd[0]
		num, _ := strconv.Atoi(cmd[1])

		switch ins {
		case "forward":
			x += num
			y += aim * num
		case "up":
			aim -= num
		case "down":
			aim += num
		default:
			fmt.Println("Error")
		}
	}

	return x * y
}
