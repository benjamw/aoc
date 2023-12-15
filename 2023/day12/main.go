package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/structs"
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
	nonos := parseInput(input)
	for _, nono := range nonos {
		_ = nono.MakeAllPossible()
	}

	n := make([]structs.Nonogram, 0)
	for _, nono := range nonos {
		n = append(n, *nono)
	}

	return algos.SumFunc(n, func(t structs.Nonogram) int {
		return t.NumPossible
	})
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	nonos := parseInput(input)
	for _, nono := range nonos {
		nono.Unfold(5)
		_ = nono.MakeAllPossible()
	}

	n := make([]structs.Nonogram, 0)
	for _, nono := range nonos {
		n = append(n, *nono)
	}

	return algos.SumFunc(n, func(t structs.Nonogram) int {
		return t.NumPossible
	})
}

func parseInput(input string) (ans []*structs.Nonogram) {
	for _, line := range strings.Split(input, "\n") {
		n := &structs.Nonogram{}

		s := strings.Split(line, " ")
		n.Line = []byte(s[0])
		n.Pattern = algos.ToIntSlice(strings.Split(s[1], ","))

		ans = append(ans, n)
	}

	return ans
}
