package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

	points := make([]int, 0)
	var same []int
	var val float64

	// find number of elements that match between the two slices
	for _, c := range parsed {
		same = algos.SimpleGeneric(c[0], c[1])
		val = math.Pow(2, float64(len(same)-1))
		points = append(points, cast.ToInt(val))
	}

	return algos.Sum(points) // 21088
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	parsed := parseInput(input)

	cards := make(map[int]int, len(parsed))

	for n := 1; n <= len(parsed); n++ {
		c := parsed[n]
		same := len(algos.SimpleGeneric(c[0], c[1]))

		if _, ok := cards[n]; !ok {
			cards[n] = 1
		}

		for nn := 1; nn <= same; nn++ {
			if _, ok := cards[n+nn]; !ok {
				cards[n+nn] = 1
			}
			cards[n+nn] += cards[n]
		}
	}

	v := make([]int, 0, len(cards))

	for _, value := range cards {
		v = append(v, value)
	}

	return algos.Sum(v) // 6874754
}

func parseInput(input string) (ans map[int][][]int) {
	ans = make(map[int][][]int)
	for _, line := range strings.Split(input, "\n") {
		// split on :
		top := strings.Split(line, ":")
		top = algos.TrimAll(top)

		reTop := regexp.MustCompile("^Card\\s*(\\d+)$")
		matched := reTop.FindStringSubmatch(top[0])
		cardId := cast.ToInt(matched[1])

		// split on |
		mid := strings.Split(top[1], "|")
		mid = algos.TrimAll(mid)

		reNum := regexp.MustCompile("\\d+")
		needS := reNum.FindAllString(mid[0], -1)
		haveS := reNum.FindAllString(mid[1], -1)

		need := make([]int, len(needS))
		have := make([]int, len(haveS))

		var i int
		var n string
		for i, n = range needS {
			need[i] = cast.ToInt(n)
		}

		for i, n = range haveS {
			have[i] = cast.ToInt(n)
		}

		ans[cardId] = [][]int{
			need,
			have,
		}
	}

	return ans
}
