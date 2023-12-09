package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/cast"
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
	parsed := parseInput(input)

	partLocs, symbLocs := getLocations(parsed, false)
	parts := lookAround(symbLocs, partLocs, false)

	return algos.Sum(parts) // 540212
}

func part2(input string) int {
	parsed := parseInput(input)

	partLocs, symbLocs := getLocations(parsed, true)
	gears := lookAround(symbLocs, partLocs, true)

	return algos.Sum(gears) // 87605697
}

func parseInput(input string) (ans map[int]string) {
	ans = make(map[int]string)
	for n, line := range strings.Split(input, "\n") {
		ans[n] = line
	}

	return ans
}

func getLocations(parsed map[int]string, doGears bool) (map[int]map[int]int, map[int]map[int]string) {
	// find the positions of each part number
	rePart := regexp.MustCompile("\\d+")
	partLocs := make(map[int]map[int]int)

	// find the index of all the symbols in each line
	reSymb := regexp.MustCompile("[^\\d.]")
	if doGears {
		reSymb = regexp.MustCompile("\\*")
	}
	symbLocs := make(map[int]map[int]string)

	for n, line := range parsed {
		partMatches := rePart.FindAllStringIndex(line, -1)
		for _, partMatch := range partMatches {
			part := cast.ToInt(line[partMatch[0]:partMatch[1]])
			for i := partMatch[0]; i < partMatch[1]; i++ {
				if partLocs[n] == nil {
					partLocs[n] = make(map[int]int)
				}
				partLocs[n][i] = part
			}
		}

		symbMatches := reSymb.FindAllStringIndex(line, -1)
		for _, symbMatch := range symbMatches {
			if symbLocs[n] == nil {
				symbLocs[n] = make(map[int]string)
			}
			symbLocs[n][symbMatch[0]] = line[symbMatch[0]:symbMatch[1]]
		}
	}

	return partLocs, symbLocs
}

func lookAround(s map[int]map[int]string, p map[int]map[int]int, doGears bool) []int {
	ret := make([]int, 0)
	var nw []int

	var part int
	var ok, pok bool

	// look at each symbol
	for y, sl := range s {
		for x, _ := range sl {
			nw = make([]int, 0)

			// check to see if there's a part in all directions around this point x,y
			// check NW (x-1, y-1)
			part, ok = p[y-1][x-1]
			if ok {
				nw = append(nw, part)
			}
			pok = ok

			// check N (x, y-1)
			part, ok = p[y-1][x]
			if ok && !pok {
				nw = append(nw, part)
			}
			pok = ok

			// check NE (x+1, y-1)
			part, ok = p[y-1][x+1]
			if ok && !pok {
				nw = append(nw, part)
			}

			// check W (x-1, y)
			part, ok = p[y][x-1]
			if ok {
				nw = append(nw, part)
			}

			// check E (x+1, y)
			part, ok = p[y][x+1]
			if ok {
				nw = append(nw, part)
			}

			// check SW (x-1, y+1)
			part, ok = p[y+1][x-1]
			if ok {
				nw = append(nw, part)
			}
			pok = ok

			// check S (x, y+1)
			part, ok = p[y+1][x]
			if ok && !pok {
				nw = append(nw, part)
			}
			pok = ok

			// check SE (x+1, y+1)
			part, ok = p[y+1][x+1]
			if ok && !pok {
				nw = append(nw, part)
			}

			if doGears {
				if 2 == len(nw) {
					ret = append(ret, nw[0]*nw[1])
				}
			} else {
				ret = append(ret, nw...)
			}
		}
	}

	return ret
}
