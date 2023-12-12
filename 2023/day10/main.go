package main

import (
	_ "embed"
	"flag"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/structs/grid"
	"github.com/benjamw/aoc/util"
)

//go:embed input.txt
var input string

var (
	DirMap = map[byte][]string{
		'-': {grid.R, grid.L},
		'|': {grid.U, grid.D},
		'L': {grid.U, grid.R},
		'F': {grid.R, grid.D},
		'7': {grid.D, grid.L},
		'J': {grid.U, grid.L},
	}
)

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

	count := followPath(g)

	return count
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	g := parseInput(input)

	fillLoop(&g)

	c := countInLoop(&g)

	// util.PrintGridFunc(g.Nodes, func(n grid.Node[byte]) string {
	// 	if n.InPipe {
	// 		return transVal(n.Value)
	// 	} else {
	// 		if n.InLoop {
	// 			return "I"
	// 		} else {
	// 			return "O"
	// 		}
	// 	}
	// })

	return c
}

func parseInput(input string) (g grid.Grid[byte]) {
	g.Nodes = make([][]grid.Node[byte], 0)
	for y, line := range strings.Split(input, "\n") {
		l := []byte(strings.TrimSpace(line))
		if x := algos.Find(l, byte('S')); -1 != x {
			g.StartX = x
			g.StartY = y
		}

		nl := make([]grid.Node[byte], 0)
		for x, b := range l {
			nl = append(nl, grid.Node[byte]{X: x, Y: y, Value: b, Grid: &g})
		}
		g.Nodes = append(g.Nodes, nl)
	}

	g.Get(g.StartX, g.StartY).IsStart = true

	return g
}

func followPath(g grid.Grid[byte]) int {
	c := 0

	r := g.GetStartNode()
	r1 := r
	r2 := r

	r1Dir, r2Dir := getStartingDirs(g)

	for {
		r1 = r1.Move(r1Dir)
		r2 = r2.Move(r2Dir)
		r1Dir = getDir(r1.Value, r1Dir)
		r2Dir = getDir(r2.Value, r2Dir)

		c++

		if reflect.DeepEqual(r1, r2) {
			break
		}
	}

	return c
}

func fillLoop(g *grid.Grid[byte]) {
	r := g.GetStartNode()
	d, _ := getStartingDirs(*g)

	// starting at the start node, proceed through the loop, setting InLoop to true,
	// until the start node is reached again (the InLoop is true)
	for {
		r.InPipe = true
		r = r.Move(d)

		if r.InPipe || r.Value == 'S' {
			break
		}

		d = getDir(r.Value, d)
	}
}

func getDir(elem byte, dir string) string {
	od := grid.Opposite[dir]
	dirs := DirMap[elem]
	if !slices.Contains(dirs, od) {
		panic(fmt.Sprintf("dir (%s) not found in (%s) %v", od, string(elem), dirs))
	}

	return algos.RemoveElem(dirs, od)[0]
}

func getStartingDirs(g grid.Grid[byte]) (s1, s2 string) {
	// look at the S position and determine which way each r is going
	// below S:
	if g.StartY < len(g.Nodes) && slices.Contains(DirMap[g.Get(g.StartX, g.StartY+1).Value], grid.U) {
		s1 = grid.D
	}
	// right of S:
	if g.StartX < len(g.Nodes[0]) && slices.Contains(DirMap[g.Get(g.StartX+1, g.StartY).Value], grid.L) {
		if "" != s1 {
			s2 = grid.R
		} else {
			s2 = grid.R
		}
	}
	// above S:
	if g.StartY > 0 && slices.Contains(DirMap[g.Get(g.StartX, g.StartY-1).Value], grid.D) {
		if "" != s1 {
			s2 = grid.U
		} else {
			s2 = grid.U
		}
	}
	// left of S:
	if g.StartX > 0 && slices.Contains(DirMap[g.Get(g.StartX-1, g.StartY).Value], grid.R) {
		if "" != s1 {
			s2 = grid.L
		} else {
			s2 = grid.L
		}
	}

	// set the "real" value of the start node
	g.Get(g.StartX, g.StartY).Value = getStartValue(s1, s2)

	return
}

func countInLoop(g *grid.Grid[byte]) int {
	c := 0
	var pv byte
	var opp bool
	for y := 0; y < len(g.Nodes); y++ {
		inLoop := false
		for x := 0; x < len(g.Nodes[y]); x++ {
			n := g.Get(x, y)
			if n.InPipe {
				opp = n.Value == 'J' && pv == 'F'
				opp = opp || n.Value == '7' && pv == 'L'
				if n.Value == '-' || opp {
					continue
				}

				inLoop = !inLoop
				pv = n.Value
			} else {
				n.InLoop = inLoop
				if inLoop {
					c++
				}
			}
		}
	}

	return c
}

func transVal(b byte) string {
	switch b {
	case '-':
		return "─"
	case '7':
		return "┐"
	case '|':
		return "│"
	case 'J':
		return "┘"
	case 'L':
		return "└"
	case 'F':
		return "┌"
	default:
		return "#"
	}
}

func getStartValue(s1, s2 string) byte {
	switch s1 {
	case grid.U:
		switch s2 {
		case grid.R:
			return 'L'
		case grid.D:
			return '|'
		case grid.L:
			return 'J'
		}
	case grid.R:
		switch s2 {
		case grid.D:
			return 'F'
		case grid.L:
			return '-'
		case grid.U:
			return 'L'
		}
	case grid.D:
		switch s2 {
		case grid.L:
			return '7'
		case grid.U:
			return '|'
		case grid.R:
			return 'F'
		}
	case grid.L:
		switch s2 {
		case grid.U:
			return 'J'
		case grid.R:
			return '-'
		case grid.D:
			return '7'
		}
	}

	return 'S'
}
