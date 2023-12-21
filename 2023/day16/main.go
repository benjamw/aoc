package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/structs/grid"
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
	laserGrid := parseInput(input)

	laserGrid.FollowLaser(grid.Split{
		Dir: grid.R,
		X:   0,
		Y:   0,
	})

	return laserGrid.CountEnergized()
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	laserGrid := parseInput(input)

	c := make([]int, 0)

	splits := laserGrid.MakeSplits()
	for _, split := range splits {
		laserGrid.Reset()
		laserGrid.FollowLaser(split)
		c = append(c, laserGrid.CountEnergized())
	}

	return algos.Max(c)
}

func parseInput(input string) *grid.LaserGrid {
	return grid.MakeGrid(input)
}
