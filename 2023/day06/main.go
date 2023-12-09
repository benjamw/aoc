package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/cast"
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

	numTimes := make([]int, 0)
	for _, v := range parsed {
		t := v[0]
		r := v[1]
		numTimes = append(numTimes, getNumTimes(t, r))
	}

	return algos.Product(numTimes)
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	parsed := parseInput(input)

	t := ""
	r := ""
	for _, v := range parsed {
		t += cast.ToString(v[0])
		r += cast.ToString(v[1])
	}

	return getNumTimes(cast.ToInt(t), cast.ToInt(r))
}

func parseInput(input string) (ans [][]int) {
	ans = make([][]int, 0)
	var time, record []string
	re := regexp.MustCompile("\\d+")
	for _, line := range strings.Split(input, "\n") {
		if time == nil {
			time = re.FindAllString(line, -1)
		} else {
			record = re.FindAllString(line, -1)
		}
	}

	for i := 0; i < len(time); i++ {
		ans = append(ans, []int{
			cast.ToInt(time[i]),
			cast.ToInt(record[i]),
		})
	}

	return ans
}

func getNumTimes(time int, record int) int {
	count := 0
	for t := 1; t < time; t++ {
		if t*(time-t) > record {
			count++
		}
	}

	return count
}
