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
		ans := part1(input, false)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string, part2 bool) int {
	defer util.Duration(util.Track("part1"))
	parsed := parseInput(input)

	extra := make([]int, 0)

	for i, p := range parsed {
		_ = i
		st := calcStack(p)
		extra = append(extra, reverseCalcStack(st, part2))
	}

	return algos.Sum(extra)
}

func part2(input string) int {
	return part1(input, true)
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " ")
		ans = append(ans, algos.ToIntSlice(l))
	}

	return ans
}

func differences(l []int) []int {
	ret := make([]int, 0)
	for i := 1; i < len(l); i++ {
		ret = append(ret, l[i]-l[i-1])
	}

	return ret
}

func calcStack(p []int) structs.Stack[[]int] {
	s := structs.Stack[[]int]{}
	s.Push(p)
	for {
		d := differences(p)
		s.Push(d)
		if isZeros(d) {
			return s
		}

		p = d
	}
}

func reverseCalcStack(st structs.Stack[[]int], front bool) int {
	// start going back up the stack and add values
	d := 0
	for !st.IsEmpty() {
		s, _ := st.Pop()

		if front {
			d = s[0] - d
		} else {
			d = d + s[len(s)-1]
		}
	}

	return d
}

func isZeros(l []int) bool {
	for _, i := range l {
		if i != 0 {
			return false
		}
	}

	return true
}
