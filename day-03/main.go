package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	power := p1(lines)
	fmt.Printf("Part one: %d\n", power)

	lsr := p2(lines)
	fmt.Printf("Part two: %d\n", lsr)
}

func p1(input []string) int64 {

	// declare some varibles
	rows := len(input)
	cols := len(input[0])
	sum := make([]int, cols)
	var dig []string
	n := 0

	for _, line := range input {
		dig = strings.Split(line, "")

		for i := 0; i < len(sum); i++ {
			n, _ = strconv.Atoi(dig[i])
			sum[i] += n
		}
	}
	fmt.Println(sum)

	var gbin, ebin string
	for _, s := range sum {
		if s > rows/2 {
			gbin += "1"
			ebin += "0"
		} else {
			gbin += "0"
			ebin += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gbin, 2, 32)
	epsilon, _ := strconv.ParseInt(ebin, 2, 32)

	return gamma * epsilon
}

func p2(data []string) int64 {
	//duplicate arrays
	ogr_data := make([]string, len(data))
	csr_data := make([]string, len(data))
	copy(ogr_data, data)
	copy(csr_data, data)

	// call rating function for each parameter
	ogr := rating("ogr", ogr_data)
	csr := rating("csr", csr_data)

	return ogr * csr
}

func rating(kind string, input []string) int64 {
	// fmt.Println(input)

	// declare some varibles
	cols := len(input[0])
	var bin, tot, rows, c int
	var keep string
	var dig []string
	var result int64

	// loop through columns
	for i := 0; i < cols; i++ {
		// reset variables
		bin, tot, rows, keep = 0, 0, 0, ""

		// sum digit of every row
		for _, line := range input {
			dig = strings.Split(line, "")
			bin, _ = strconv.Atoi(dig[i])
			tot += bin
		}

		rows = len(input)
		// fmt.Println(tot, rows/2)
		if kind == "ogr" {
			if tot >= rows-tot {
				keep = "1"
			} else {
				keep = "0"
			}
		} else if kind == "csr" {
			if tot >= rows-tot {
				keep = "0"
			} else {
				keep = "1"
			}
		}

		c = 0
		for _, line := range input {
			dig = strings.Split(line, "")
			if dig[i] == keep {
				input[c] = line
				c++
			}
		}
		input = input[:c]

		// break loop if single slice
		if len(input) < 2 {
			break
		}

		// fmt.Println(input)

	}

	result, _ = strconv.ParseInt(input[0], 2, 32)

	return result

}
