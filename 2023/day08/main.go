package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/structs"
	"github.com/benjamw/aoc/util"
)

//go:embed input.txt
var input string

var nodes map[string]*structs.Node

var nodeRegex *regexp.Regexp

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	nodeRegex = regexp.MustCompile("(\\w{3}) = \\((\\w{3}), (\\w{3})\\)")
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
	desert := parseInput(input)

	return desert.WalkToEnd(false)
}

func part2(input string) int {
	defer util.Duration(util.Track("part2"))
	desert := parseInput(input)

	startingNodes := getStartingNodes(nodes)

	steps := make([]int, 0)
	for _, sn := range startingNodes {
		desert.Network.Root = sn
		steps = append(steps, desert.WalkToEnd(true))
	}

	// find the LCD of the steps and find the
	return algos.LCMs(steps)
}

func parseInput(input string) (ans desert) {
	nodes = make(map[string]*structs.Node)
	desert := desert{}
	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			desert.Directions = line
			continue
		} else if "" == line {
			continue
		}

		generateNode(line)
	}

	desert.Network = structs.Network{
		Root: nodes["AAA"],
	}

	return desert
}

func generateNode(line string) {
	m := nodeRegex.FindStringSubmatch(line)
	name := m[1]
	left := m[2]
	right := m[3]

	if _, ok := nodes[name]; !ok {
		nodes[name] = &structs.Node{
			Self:  name,
			Left:  nil,
			Right: nil,
		}
	}

	if _, ok := nodes[left]; !ok {
		nodes[left] = &structs.Node{
			Self:  left,
			Left:  nil,
			Right: nil,
		}
	}
	nodes[name].Left = nodes[left]

	if _, ok := nodes[right]; !ok {
		nodes[right] = &structs.Node{
			Self:  right,
			Left:  nil,
			Right: nil,
		}
	}
	nodes[name].Right = nodes[right]
}

func getStartingNodes(n map[string]*structs.Node) []*structs.Node {
	ret := make([]*structs.Node, 0)
	for k, v := range n {
		if k[len(k)-1:] == "A" {
			ret = append(ret, v)
		}
	}

	return ret
}

type desert struct {
	Directions string
	Network    structs.Network
}

func (d desert) WalkToEnd(part2 bool) int {
	steps := 0

	n := d.Network.Root
	dirs := []rune(d.Directions)
	di := 0

	for {
		switch dirs[di] {
		case 'R':
			n = n.GoRight()
		case 'L':
			n = n.GoLeft()
		}

		steps += 1
		di += 1
		di = di % len(dirs)

		if !part2 && n.Self == "ZZZ" {
			break
		} else if part2 && n.Self[len(n.Self)-1:] == "Z" {
			break
		}
	}

	return steps
}

func (d desert) SimultWalkToEnd(startingNodes []*structs.Node) int {
	steps := 0

	ns := startingNodes
	ln := len(ns)
	dirs := []rune(d.Directions)
	di := 0

	for {
		switch dirs[di] {
		case 'R':
			for i := 0; i < ln; i++ {
				ns[i] = ns[i].GoRight()
			}
		case 'L':
			for i := 0; i < ln; i++ {
				ns[i] = ns[i].GoLeft()
			}
		}

		steps += 1
		di += 1
		di = di % len(dirs)

		nonZFound := false
		for i := 0; i < ln; i++ {
			nonZFound = nonZFound || ns[i].Self[len(ns[i].Self)-1:] != "Z"
			if nonZFound {
				break
			}
		}

		if !nonZFound {
			break
		}
	}

	return steps
}
