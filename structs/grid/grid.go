package grid

import (
	"errors"
	"fmt"
)

const (
	U = "u"
	R = "r"
	D = "d"
	L = "l"
	N = "u"
	E = "r"
	S = "d"
	W = "l"
)

var (
	Opposite = map[string]string{
		U: D,
		R: L,
		D: U,
		L: R,
	}
)

type Grid[T any] struct {
	Nodes  [][]Node[T]
	StartX int
	StartY int
}

func (g *Grid[T]) Get(x, y int) *Node[T] {
	return &g.Nodes[y][x]
}

func (g *Grid[T]) GetStartNode() *Node[T] {
	return g.Get(g.StartX, g.StartY)
}

type Node[T any] struct {
	X       int
	Y       int
	Value   T
	InPipe  bool
	Grid    *Grid[T]
	InLoop  bool
	IsStart bool
}

func (n *Node[T]) Up() (nn *Node[T], err error) {
	if n.Y == 0 {
		err = errors.New("y out of bounds low")
		return
	}

	nn = &n.Grid.Nodes[n.Y-1][n.X]
	return
}

func (n *Node[T]) Right() (nn *Node[T], err error) {
	if n.X+1 >= len(n.Grid.Nodes[0]) {
		err = errors.New("x out of bounds high")
		return
	}

	nn = &n.Grid.Nodes[n.Y][n.X+1]
	return
}

func (n *Node[T]) Down() (nn *Node[T], err error) {
	if n.Y+1 >= len(n.Grid.Nodes) {
		err = errors.New("y out of bounds high")
		return
	}

	nn = &n.Grid.Nodes[n.Y+1][n.X]
	return
}

func (n *Node[T]) Left() (nn *Node[T], err error) {
	if n.X-1 < 0 {
		err = errors.New("x out of bounds low")
		return
	}

	nn = &n.Grid.Nodes[n.Y][n.X-1]
	return
}

func (n *Node[T]) Move(dir string) *Node[T] {
	var nx, ny int
	var err error
	var nn *Node[T]

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

	if err != nil {
		panic(fmt.Sprintf("outside grid @ (%d, %d): %s", nx, ny, err.Error()))
	}

	return nn
}
