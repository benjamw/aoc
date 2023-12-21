package main

import (
	_ "embed"
	"flag"
	"fmt"
	"reflect"
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
	return doThing(input, false)
}

func part2(input string) int {
	return doThing(input, true)
}

func doThing(input string, smudge bool) int {
	defer util.Duration(util.Track("part1"))
	parsed := parseInput(input)

	c := 0

	for i, pod := range parsed {
		above, found, fixed := findReflection(pod, smudge)
		if found {
			if !smudge || (smudge && fixed) {
				c += above * 100
				continue
			}
		}

		pod = algos.Transpose(pod)
		left, found, fixed := findReflection(pod, smudge)
		if !found {
			panic(fmt.Sprintf("not found after both passes on pod %d", i))
		}

		c += left
	}

	return c
}

func parseInput(input string) (ans [][][]byte) {
	for _, pods := range strings.Split(input, "\n\n") {
		pod := make([][]byte, 0)
		for _, line := range strings.Split(pods, "\n") {
			pod = append(pod, []byte(line))
		}

		ans = append(ans, pod)
	}

	return ans
}

func findReflection(pod [][]byte, smudge bool) (above int, found bool, fixed bool) {
	found = false
	fixed = false
	var l []byte
	var m []byte
	var f bool
	numFixed := 0

	// find a spot where this line matches the next line
	for i := 0; i < len(pod)-1; i++ {
		numFixed = 0
		fixed = false

		if smudge {
			// fix any single character differences
			l, f = fixSmudge(pod[i], pod[i+1])
			if f {
				fixed = fixed || f
				numFixed++
			}
		}

		if reflect.DeepEqual(l, pod[i+1]) {
			found = true
			above = i + 1
			// check to edges
			for j := i; j >= 0; j-- {
				if i-j+above >= len(pod) {
					break
				}

				m = pod[j]
				if smudge {
					// fix any single character differences
					m, f = fixSmudge(pod[j], pod[i-j+above])
					if j != i && f {
						fixed = fixed || f
						numFixed++
					}
				}

				// if there were too many lines fixed
				if numFixed > 1 {
					found = false
					break
				}

				// if it's not outside the grid, and the line doesn't match
				if !reflect.DeepEqual(m, pod[i-j+above]) {
					found = false
					break
				}
			}

			if (!smudge && found) || (smudge && fixed && found) {
				break
			}
		}
	}

	return
}

func fixSmudge(line1 []byte, line2 []byte) (line []byte, fixed bool) {
	c := 0
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			c++
		}
	}

	if c != 1 {
		// return the original line
		return line1, false
	}

	// return the "fixed" line
	return line2, true
}
