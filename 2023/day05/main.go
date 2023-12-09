package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/util"
)

//go:embed input.txt
var input string

type garden struct {
	seeds        []int
	seedToSoil   [][]int
	soilToFert   [][]int
	fertToWater  [][]int
	waterToLight [][]int
	lightToTemp  [][]int
	tempToHumid  [][]int
	humidToLoc   [][]int
}

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
	g := parseInput(input)

	// check each seed for the location it matches to
	l := make([]int, 0)
	for _, s := range g.seeds {
		l = append(l, seedToLocation(s, g))
	}

	return slices.Min(l)
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	g := parseInput(input)

	// check each seed range for the location it matches to
	l := make([]int, 0)
	for n := 0; n < len(g.seeds); n += 2 {
		lt := make([]int, 0)
		for nn := 0; nn < g.seeds[n+1]; nn++ {
			lt = append(lt, seedToLocation(g.seeds[n]+nn, g))
		}
		l = append(l, slices.Min(lt))
	}

	return slices.Min(l)
}

func parseInput(input string) (ans garden) {
	var s2s, s2f, f2w, w2l, l2t, t2h, h2l bool
	r := regexp.MustCompile("\\d+")
	ret := garden{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "seeds") {
			seeds := r.FindAllString(line, -1)
			ret.seeds = algos.ToIntSlice(seeds)
			continue
		}

		if line == "" {
			s2s, s2f, f2w, w2l, l2t, t2h, h2l = false, false, false, false, false, false, false
			continue
		}

		switch line {
		case "seed-to-soil map:":
			s2s = true
			continue
		case "soil-to-fertilizer map:":
			s2f = true
			continue
		case "fertilizer-to-water map:":
			f2w = true
			continue
		case "water-to-light map:":
			w2l = true
			continue
		case "light-to-temperature map:":
			l2t = true
			continue
		case "temperature-to-humidity map:":
			t2h = true
			continue
		case "humidity-to-location map:":
			h2l = true
			continue
		}

		switch true {
		case s2s:
			ret.seedToSoil = processMap(ret.seedToSoil, r.FindAllString(line, -1))
		case s2f:
			ret.soilToFert = processMap(ret.soilToFert, r.FindAllString(line, -1))
		case f2w:
			ret.fertToWater = processMap(ret.fertToWater, r.FindAllString(line, -1))
		case w2l:
			ret.waterToLight = processMap(ret.waterToLight, r.FindAllString(line, -1))
		case l2t:
			ret.lightToTemp = processMap(ret.lightToTemp, r.FindAllString(line, -1))
		case t2h:
			ret.tempToHumid = processMap(ret.tempToHumid, r.FindAllString(line, -1))
		case h2l:
			ret.humidToLoc = processMap(ret.humidToLoc, r.FindAllString(line, -1))
		}
	}

	return ret
}

func processMap(m [][]int, l []string) [][]int {
	s := algos.ToIntSlice(l)

	if m == nil {
		m = make([][]int, 0)
	}
	m = append(m, s)

	return m
}

func seedToLocation(s int, g garden) int {
	so := translate(s, g.seedToSoil)
	f := translate(so, g.soilToFert)
	w := translate(f, g.fertToWater)
	li := translate(w, g.waterToLight)
	t := translate(li, g.lightToTemp)
	h := translate(t, g.tempToHumid)
	return translate(h, g.humidToLoc)
}

func translate(v int, m [][]int) int {
	// for each map
	for _, mm := range m {
		d, s, r := mm[0], mm[1], mm[2]
		if v >= s && v < s+r { // if the value is in the given source range
			return d + (v - s) // return the destination value plus the count difference
		}
	}

	// no match found, just return the number
	return v
}
