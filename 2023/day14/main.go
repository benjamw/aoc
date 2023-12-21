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
	defer util.Duration(util.Track("part1"))
	grid := parseInput(input)

	grid = algos.MoveNorth(grid, ".", "#", "O")

	return sumStones(grid)
}

func part2(input string) int {
	defer util.Duration(util.Track("part1"))
	grid := parseInput(input)

	var prevGrid [][]string
	for {
		prevGrid = algos.Clone(grid)
		grid = processSpin(grid, 1e9)

		if reflect.DeepEqual(grid, prevGrid) {
			break
		}
	}

	return sumStones(grid)
}

func processSpin(grid [][]string, cycles int) [][]string {
	for i := 0; i < cycles; i++ {
		grid = algos.MoveNorth(grid, ".", "#", "O")
		grid = algos.MoveWest(grid, ".", "#", "O")
		grid = algos.MoveSouth(grid, ".", "#", "O")
		grid = algos.MoveEast(grid, ".", "#", "O")
	}

	return grid
}

func parseInput(input string) (ans [][]string) {
	return algos.MakeGrid(input)
}

func sumStones(grid [][]string) int {
	t := 0
	for i := 0; i < len(grid); i++ {
		score := len(grid) - i

		// count the number of Os on this level
		c := 0
		for _, v := range grid[i] {
			if v == "O" {
				c++
			}
		}

		t += score * c
	}

	return t
}
