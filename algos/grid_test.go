package algos

import (
	"reflect"
	"testing"
)

var in = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

var outN = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

var outS = `.....#....
....#....#
...O.##...
...#......
O.O....O#O
O.#..O.#.#
O....#....
OO....OO..
#OO..###..
#OO.O#...O`

var outE = `....O#....
.OOO#....#
.....##...
.OO#....OO
......OO#.
.O#...O#.#
....O#..OO
.........O
#....###..
#..OO#....`

var outW = `O....#....
OOO.#....#
.....##...
OO.#OO....
OO......#.
O.#O...#.#
O....#OO..
O.........
#....###..
#OO..#....`

func TestMoveNorth(t *testing.T) {
	inGrid := MakeGrid(in)
	outGrid := MakeGrid(outN)

	type args[T comparable] struct {
		grid      [][]T
		space     T
		immovable T
		movable   T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[string]{
		{
			name: "main",
			args: args[string]{
				grid:      inGrid,
				space:     ".",
				immovable: "#",
				movable:   "O",
			},
			want: outGrid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveNorth(tt.args.grid, tt.args.space, tt.args.immovable, tt.args.movable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveNorth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveSouth(t *testing.T) {
	inGrid := MakeGrid(in)
	outGrid := MakeGrid(outS)

	type args[T comparable] struct {
		grid      [][]T
		space     T
		immovable T
		movable   T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[string]{
		{
			name: "main",
			args: args[string]{
				grid:      inGrid,
				space:     ".",
				immovable: "#",
				movable:   "O",
			},
			want: outGrid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveSouth(tt.args.grid, tt.args.space, tt.args.immovable, tt.args.movable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveNorth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveEast(t *testing.T) {
	inGrid := MakeGrid(in)
	outGrid := MakeGrid(outE)

	type args[T comparable] struct {
		grid      [][]T
		space     T
		immovable T
		movable   T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[string]{
		{
			name: "main",
			args: args[string]{
				grid:      inGrid,
				space:     ".",
				immovable: "#",
				movable:   "O",
			},
			want: outGrid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveEast(tt.args.grid, tt.args.space, tt.args.immovable, tt.args.movable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveNorth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveWest(t *testing.T) {
	inGrid := MakeGrid(in)
	outGrid := MakeGrid(outW)

	type args[T comparable] struct {
		grid      [][]T
		space     T
		immovable T
		movable   T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[string]{
		{
			name: "main",
			args: args[string]{
				grid:      inGrid,
				space:     ".",
				immovable: "#",
				movable:   "O",
			},
			want: outGrid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveWest(tt.args.grid, tt.args.space, tt.args.immovable, tt.args.movable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveNorth() = %v, want %v", got, tt.want)
			}
		})
	}
}
