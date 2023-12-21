package grid

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/benjamw/aoc/algos"
)

type Split struct {
	Dir string
	X   int
	Y   int
}

type LaserGrid struct {
	Nodes    [][]*LaserNode
	curNode  *LaserNode
	splits   []Split
	curSplit Split
	dir      string
	x        int
	y        int
}

func (g *LaserGrid) Get(x, y int) (*LaserNode, error) {
	if y >= len(g.Nodes) || y < 0 {
		return nil, errors.New("y out of bounds")
	}

	if x >= len(g.Nodes[0]) || x < 0 {
		return nil, errors.New("x out of bounds")
	}

	return g.Nodes[y][x], nil
}

func (g *LaserGrid) FollowLaser(start Split) {
	var dir string

	g.splits = []Split{start}
	g.curSplit = g.splits[0]
	g.dir = g.curSplit.Dir
	g.x = g.curSplit.X
	g.y = g.curSplit.Y
	g.curNode, _ = g.Get(g.x, g.y)

	for len(g.splits) > 0 {
		// check if this node has already been traversed in this direction...
		if g.curNode.Energized && slices.Contains(g.curNode.DirsTraversed, g.dir) {
			g.getNextSplit()
			continue
		}

		g.curNode.Energized = true
		g.curNode.DirsTraversed = append(g.curNode.DirsTraversed, g.dir)

		switch g.curNode.Value {
		case "|":
			if g.dir == U || g.dir == D {
				// pass right through
				g.doMove(g.dir)
				continue
			} else {
				// add a split for the down direction
				g.splits = append(g.splits, Split{
					Dir: D,
					X:   g.x,
					Y:   g.y,
				})

				// continue on in the up direction
				g.doMove(U)
				continue
			}
		case "-":
			if g.dir == L || g.dir == R {
				// pass right through
				g.doMove(g.dir)
				continue
			} else {
				// add a split for the left direction
				g.splits = append(g.splits, Split{
					Dir: L,
					X:   g.x,
					Y:   g.y,
				})

				// continue on in the right direction
				g.doMove(R)
				continue
			}
		case "\\":
			switch g.dir {
			case U:
				dir = L
			case R:
				dir = D
			case D:
				dir = R
			case L:
				dir = U
			}
			g.doMove(dir)
			continue
		case "/":
			switch g.dir {
			case U:
				dir = R
			case R:
				dir = U
			case D:
				dir = L
			case L:
				dir = D
			}
			g.doMove(dir)
			continue
		default: // "."
			g.doMove(g.dir)
			continue
		}
	}
}

func (g *LaserGrid) doMove(dir string) {
	nextNode, err := g.curNode.Move(dir)
	if err != nil {
		// this path has gone off the grid
		g.getNextSplit()
		return
	}

	g.curNode = nextNode
	g.dir = dir
	g.x = g.curNode.X
	g.y = g.curNode.Y
}

func (g *LaserGrid) getNextSplit() {
	var err error
	g.splits = algos.RemoveElem(g.splits, g.curSplit)

	if len(g.splits) > 0 {
		g.curSplit = g.splits[0]
		g.dir = g.curSplit.Dir
		g.x = g.curSplit.X
		g.y = g.curSplit.Y
		g.curNode, err = g.Get(g.x, g.y)
		if err != nil {
			panic("split contained invalid node")
		}
	}
}

func (g *LaserGrid) CountEnergized() int {
	c := 0
	for _, row := range g.Nodes {
		for _, node := range row {
			if node.Energized {
				c++
			}
		}
	}

	return c
}

func (g *LaserGrid) PrintGrid() {
	fmt.Printf("\n\n")
	for _, row := range g.Nodes {
		for _, node := range row {
			fmt.Printf("%s ", node.Value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func (g *LaserGrid) PrintEnergized() {
	fmt.Printf("\n\n")
	for _, row := range g.Nodes {
		for _, node := range row {
			out := "."
			if node.Energized {
				out = "#"
			}
			fmt.Printf("%s ", out)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func (g *LaserGrid) MakeSplits() []Split {
	s := make([]Split, 0)

	// create all the top and bottom splits
	for i := 0; i < len(g.Nodes[0]); i++ {
		s = append(s, Split{
			Dir: D,
			X:   i,
			Y:   0,
		})

		s = append(s, Split{
			Dir: U,
			X:   i,
			Y:   len(g.Nodes) - 1,
		})
	}

	// create all the left and right splits
	for i := 0; i < len(g.Nodes); i++ {
		s = append(s, Split{
			Dir: R,
			X:   0,
			Y:   i,
		})

		s = append(s, Split{
			Dir: L,
			X:   len(g.Nodes[0]) - 1,
			Y:   i,
		})
	}

	return s
}

func (g *LaserGrid) Reset() {
	// reset all the nodes
	for _, row := range g.Nodes {
		for _, node := range row {
			node.Reset()
		}
	}
}

type LaserNode struct {
	Grid          *LaserGrid
	Energized     bool
	Value         string
	X             int
	Y             int
	DirsTraversed []string
}

func (n *LaserNode) Reset() {
	n.DirsTraversed = make([]string, 0)
	n.Energized = false
}

func (n *LaserNode) Up() (nn *LaserNode, err error) {
	if n.Y == 0 {
		err = errors.New("y out of bounds low")
		return
	}

	return n.Grid.Get(n.X, n.Y-1)
}

func (n *LaserNode) Right() (nn *LaserNode, err error) {
	if n.X+1 >= len(n.Grid.Nodes[0]) {
		err = errors.New("x out of bounds high")
		return
	}

	return n.Grid.Get(n.X+1, n.Y)
}

func (n *LaserNode) Down() (nn *LaserNode, err error) {
	if n.Y+1 >= len(n.Grid.Nodes) {
		err = errors.New("y out of bounds high")
		return
	}

	return n.Grid.Get(n.X, n.Y+1)
}

func (n *LaserNode) Left() (nn *LaserNode, err error) {
	if n.X-1 < 0 {
		err = errors.New("x out of bounds low")
		return
	}

	return n.Grid.Get(n.X-1, n.Y)
}

func (n *LaserNode) Move(dir string) (nn *LaserNode, err error) {
	switch dir {
	case U:
		nn, err = n.Up()
	case R:
		nn, err = n.Right()
	case D:
		nn, err = n.Down()
	case L:
		nn, err = n.Left()
	}

	return nn, err
}

func MakeGrid(input string) *LaserGrid {
	grid := &LaserGrid{}

	nodes := make([][]*LaserNode, 0)
	for y, line := range strings.Split(input, "\n") {
		lineNodes := make([]*LaserNode, 0)
		for x, v := range strings.Split(line, "") {
			node := &LaserNode{
				Grid:          grid,
				Energized:     false,
				Value:         v,
				X:             x,
				Y:             y,
				DirsTraversed: make([]string, 0),
			}

			lineNodes = append(lineNodes, node)
		}
		nodes = append(nodes, lineNodes)
	}

	grid.Nodes = nodes

	return grid
}
