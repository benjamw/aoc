package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	defer util.Duration(util.Track("part1"))
	parsed := parseInput(input)

	sum := 0
	for _, v := range parsed {
		if isSafe(v) {
			sum++
		}
	}

	return sum
}

func part2(input string) int {
	defer util.Duration(util.Track("part1"))
	parsed := parseInput(input)

	sum := 0
outerloop: // oh, dear christ... they've come for the children!
	for _, v := range parsed {
		if isSafe(v) {
			sum++
			continue
		}

		// I am the brute squad
		for i := 0; i < len(v); i++ {
			v2 := algos.RemoveIndex(v, i)
			if isSafe(v2) {
				sum++
				continue outerloop
			}
		}
	}

	return sum
}

func parseInput(input string) (ans [][]int) {
	slice := algos.ToMultiSlice(input)
	ans = make([][]int, len(slice))
	for i, s := range slice {
		ans[i] = algos.ToIntSlice(s)
	}

	return
}

func isSafe(s []int) bool {
	increasing := false
	for i, v := range s {
		if i+1 >= len(s) {
			break
		}

		diff := v - s[i+1]
		if i == 0 && diff > 0 {
			increasing = true
		}

		if (increasing && diff < 0) || (!increasing && diff > 0) || (diff == 0) {
			return false
		}

		diff = int(math.Abs(float64(diff)))
		if diff > 3 {
			return false
		}
	}

	return true
}
