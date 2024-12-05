package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
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
	for i, v := range parsed[0] {
		sum += int(math.Abs(float64(v - parsed[1][i])))
	}

	return sum
}

func part2(input string) int {
	defer util.Duration(util.Track("part1"))
	parsed := parseInput(input)

	sum := 0
	for _, v := range parsed[0] {
		count := algos.Count(parsed[1], func(x int) bool { return x == v })
		sum += v * count
	}

	return sum
}

func parseInput(input string) (ans [][]int) {
	slice := algos.ToColumnSlice(input)
	cols := make([][]int, len(slice))
	for i, s := range slice {
		cols[i] = algos.ToIntSlice(s)
	}

	for _, col := range cols {
		// `col` is a reference, so no saving necessary
		sort.Ints(col)
	}

	return cols
}
