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

type PipeGrid[T any] struct {
	Nodes  [][]PipeNode[T]
	StartX int
	StartY int
}

func (g *PipeGrid[T]) Get(x, y int) *PipeNode[T] {
	return &g.Nodes[y][x]
}

func (g *PipeGrid[T]) GetStartNode() *PipeNode[T] {
	return g.Get(g.StartX, g.StartY)
}

type PipeNode[T any] struct {
	X       int
	Y       int
	Value   T
	InPipe  bool
	Grid    *PipeGrid[T]
	InLoop  bool
	IsStart bool
}

func (n *PipeNode[T]) Up() (nn *PipeNode[T], err error) {
	if n.Y == 0 {
		err = errors.New("y out of bounds low")
		return
	}

	nn = &n.Grid.Nodes[n.Y-1][n.X]
	return
}

func (n *PipeNode[T]) Right() (nn *PipeNode[T], err error) {
	if n.X+1 >= len(n.Grid.Nodes[0]) {
		err = errors.New("x out of bounds high")
		return
	}

	nn = &n.Grid.Nodes[n.Y][n.X+1]
	return
}

func (n *PipeNode[T]) Down() (nn *PipeNode[T], err error) {
	if n.Y+1 >= len(n.Grid.Nodes) {
		err = errors.New("y out of bounds high")
		return
	}

	nn = &n.Grid.Nodes[n.Y+1][n.X]
	return
}

func (n *PipeNode[T]) Left() (nn *PipeNode[T], err error) {
	if n.X-1 < 0 {
		err = errors.New("x out of bounds low")
		return
	}

	nn = &n.Grid.Nodes[n.Y][n.X-1]
	return
}

func (n *PipeNode[T]) Move(dir string) *PipeNode[T] {
	var nx, ny int
	var err error
	var nn *PipeNode[T]

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
