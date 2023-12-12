package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"slices"
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
		ans := part2(input, 1e6)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	exp := 2
	defer util.Duration(util.Track("part1"))
	space := parseInput(input)

	var expRows, expCols []int
	space, expRows, expCols = expand(space)
	galaxies := findGalaxies(space)

	distances := make([]int, 0)
	for i, g1 := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			distances = append(distances, dist(g1[0], g1[1], g2[0], g2[1], expRows, expCols, exp))
		}
	}

	return algos.Sum(distances)
}

func part2(input string, exp int) int {
	defer util.Duration(util.Track("part2"))
	space := parseInput(input)

	var expRows, expCols []int
	space, expRows, expCols = expand(space)
	galaxies := findGalaxies(space)

	distances := make([]int, 0)
	for i, g1 := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			distances = append(distances, dist(g1[0], g1[1], g2[0], g2[1], expRows, expCols, exp))
		}
	}

	return algos.Sum(distances)
}

func parseInput(input string) [][]byte {
	grid := make([][]byte, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []byte(strings.TrimSpace(line)))
	}

	return grid
}

func expand(g [][]byte) ([][]byte, []int, []int) {
	var expRows, expCols []int
	g, expRows = expandSpace(g)

	// transpose so the process is simpler for the columns
	g = algos.Transpose(g)
	g, expCols = expandSpace(g)

	// transpose back to original shape
	g = algos.Transpose(g)

	return g, expRows, expCols
}

func expandSpace(grid [][]byte) ([][]byte, []int) {
	er := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		row := grid[i]
		if !bytes.Contains(row, []byte("#")) {
			er = append(er, i)
			grid = slices.Replace(grid, i, i+1, []byte(strings.Repeat("E", len(grid[i]))))
		}
	}

	return grid, er
}

func findGalaxies(grid [][]byte) [][]int {
	l := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if '#' == grid[i][j] {
				l = append(l, []int{i, j})
			}
		}
	}

	return l
}

func dist(sx, sy, ex, ey int, er, ec []int, f int) int {
	f = f - 1 // subtract 1, because 1 was already there before expanding

	br := getExpBetween(er, sx, ex)
	bc := getExpBetween(ec, sy, ey)
	nDist := int(math.Abs(float64(sx)-float64(ex)) + math.Abs(float64(sy)-float64(ey)))
	return nDist + (f * br) + (f * bc)
}

func getExpBetween(s []int, a, b int) int {
	c := 0
	u := a < b
	for _, v := range s {
		if u {
			if a < v && v < b {
				c++
			}
		} else {
			if b < v && v < a {
				c++
			}
		}
	}

	return c
}
