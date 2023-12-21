package main

import (
	"reflect"
	"testing"

	"github.com/benjamw/aoc/algos"
)

var example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  136,
		},
		{
			name:  "actual",
			input: input,
			want:  106990,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  64,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

var want1 = `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`

var want2 = `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`

var want3 = `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`

func Test_processSpin(t *testing.T) {
	type args struct {
		grid   [][]string
		cycles int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "first",
			args: args{
				grid:   algos.MakeGrid(example),
				cycles: 1,
			},
			want: algos.MakeGrid(want1),
		},
		{
			name: "second",
			args: args{
				grid:   algos.MakeGrid(example),
				cycles: 2,
			},
			want: algos.MakeGrid(want2),
		},
		{
			name: "third",
			args: args{
				grid:   algos.MakeGrid(example),
				cycles: 3,
			},
			want: algos.MakeGrid(want3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processSpin(tt.args.grid, tt.args.cycles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processSpin() = %v, want %v", got, tt.want)
			}
		})
	}
}
