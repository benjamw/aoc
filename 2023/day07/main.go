package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/cast"
	"github.com/benjamw/aoc/util"
	"github.com/davecgh/go-spew/spew"
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
	parsed := parseInput(input, false)

	winnings := make([]int, 0)
	for i, v := range parsed {
		winnings = append(winnings, (i+1)*v.Value)
	}

	return algos.Sum(winnings)
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	parsed := parseInput(input, true)

	winnings := make([]int, 0)
	for i, v := range parsed {
		winnings = append(winnings, (i+1)*v.Value)
	}

	return algos.Sum(winnings)
}

func parseInput(input string, jokers bool) []hand {
	h := make([]hand, 0)
	re := regexp.MustCompile("\\S+")
	for _, line := range strings.Split(input, "\n") {
		e := re.FindAllString(line, -1)
		h = append(h, hand{e[0], charFrequency(e[0], jokers), strength(e[0], jokers), cast.ToInt(e[1])})
	}

	sort.Slice(h, func(i, j int) bool {
		return sortHands(h[i], h[j])
	})

	k := make([]string, 0)
	for _, h2 := range h {
		fc := make([]int, 0)
		for _, i2 := range h2.Freq {
			fc = append(fc, i2.Count)
		}
		if len(fc) == 1 {
			fc = append(fc, 0)
		}
		k = append(k, fmt.Sprintf("%s: %s: %s",
			h2.Key, strings.Join(algos.ToStringSlice(fc[:2]), "-"), strings.Join(algos.ToStringSlice(h2.Strength), "-")))
	}
	spew.Dump(k)

	return h
}

func sortHands(a, b hand) bool {
	fa := a.Freq
	fb := b.Freq
	la := len(fa)
	lb := len(fb)
	l := min(la, lb)
	for i := 0; i < l; i++ {
		if fa[i].Count == fb[i].Count {
			continue
		} else {
			return fa[i].Count < fb[i].Count
		}
	}

	ra := a.Strength
	rb := b.Strength
	for i := 0; i < len(ra); i++ {
		if ra[i] == rb[i] {
			continue
		} else {
			return ra[i] < rb[i]
		}
	}

	return false
}

func charFrequency(s string, jokers bool) []frequency {
	f := make(map[rune]int)
	for _, char := range []rune(s) {
		if _, ok := f[char]; ok {
			continue
		}

		f[char] = strings.Count(s, string(char))
	}

	fm := make([]frequency, 0)
	for r, c := range f {
		fm = append(fm, frequency{r, c})
	}

	sort.Slice(fm, func(i, j int) bool {
		// greater than to sort in desc order
		return fm[i].Count > fm[j].Count
	})

	if jokers {
		jStored := 0
		for i, v := range fm {
			if 'J' == v.Char {
				if fm[i].Count == 5 {
					continue
				}

				if i != 0 {
					fm[0].Count += v.Count
				} else {
					jStored = v.Count
				}

				fm[i].Count = 0
			} else if jStored > 0 {
				fm[i].Count += jStored
				jStored = 0
			}
		}

		// sort them again
		sort.Slice(fm, func(i, j int) bool {
			// greater than to sort in desc order
			return fm[i].Count > fm[j].Count
		})
	}

	return fm
}

func strength(s string, jokers bool) []int {
	sv := []rune(s)
	iv := make([]int, len(sv))
	for i, r := range sv {
		switch r {
		case 'A':
			iv[i] = 14
		case 'K':
			iv[i] = 13
		case 'Q':
			iv[i] = 12
		case 'J':
			if jokers {
				iv[i] = 1
			} else {
				iv[i] = 11
			}
		case 'T':
			iv[i] = 10
		default:
			iv[i] = cast.ToInt(r)
		}
	}

	return iv
}

type frequency struct {
	Char  rune
	Count int
}
type hand struct {
	Key      string
	Freq     []frequency
	Strength []int
	Value    int
}
