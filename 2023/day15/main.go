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

	c := 0
	for _, bytes := range parsed {
		c += hash(bytes)
	}

	return c
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	parsed := parseInput(input)

	reDash := regexp.MustCompile("^(\\w+)-$")
	reEqual := regexp.MustCompile("^(\\w+)=(\\d+)$")

	boxes := make([]*Box, 256)
	lenses := make(map[string]*Lens, 0)

	for _, bytes := range parsed {
		s := string(bytes)

		if strings.Contains(s, "=") {
			m := reEqual.FindStringSubmatch(s)
			bn := hash([]byte(m[1]))

			if boxes[bn] == nil {
				boxes[bn] = &Box{
					position: bn,
					lenses:   make([]*Lens, 0),
				}
			}

			l := &Lens{
				label: m[1],
				focal: cast.ToInt(m[2]),
			}
			lenses[(*l).label] = l
			boxes[bn].PlaceLens(l)
		} else if strings.Contains(s, "-") {
			m := reDash.FindStringSubmatch(s)
			bn := hash([]byte(m[1]))

			if boxes[bn] == nil {
				boxes[bn] = &Box{
					position: bn,
					lenses:   make([]*Lens, 0),
				}
			}

			boxes[bn].RemoveLensByLabel(m[1])
		}
	}

	powers := make([]int, 0)
	for i, box := range boxes {
		if box == nil || len((*box).lenses) == 0 {
			continue
		}

		powers = append(powers, getPower(i, (*box).lenses))
	}

	return algos.Sum(powers)
}

func parseInput(input string) (ans [][]byte) {
	for _, line := range strings.Split(input, ",") {
		ans = append(ans, []byte(line))
	}

	return ans
}

func hash(chars []byte) (h int) {
	h = 0
	for _, char := range chars {
		h += int(char)
		h *= 17
		h %= 256
	}

	return
}

func getPower(box int, lenses []*Lens) int {
	p := make([]int, 0)
	for i, lens := range lenses {
		pl := box + 1
		pl *= i + 1
		pl *= (*lens).focal

		p = append(p, pl)
	}

	return algos.Sum(p)
}

type Lens struct {
	label string
	focal int
	box   *Box
}

func (l *Lens) Remove() {
	(*l).box.RemoveLens(l)
	(*l).box = nil
}

type Box struct {
	position int
	lenses   []*Lens
}

func (b *Box) PlaceLens(l *Lens) {
	(*l).box = b

	// search the box for a current lens with the same label
	for i, lens := range (*b).lenses {
		if (*lens).label == (*l).label {
			(*b).lenses[i] = l
			lens.box = nil

			return
		}
	}

	// if not found, place at the end of the box
	(*b).lenses = append((*b).lenses, l)
}

func (b *Box) RemoveLensByLabel(l string) {
	for _, lens := range (*b).lenses {
		if (*lens).label == l {
			b.RemoveLens(lens)
		}
	}
}

func (b *Box) RemoveLens(l *Lens) {
	(*l).box = nil
	(*b).lenses = algos.RemoveElem((*b).lenses, l)
}
