package main

import (
	"testing"
)

var example = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

var example2 = `...#...##...###
##.#.##.#..#.##
####.##.#..#.##
#.#.##.###.#.##
#...##.##...#.#
##..#.###...#..
#...####.###...
#.#####...##.##
#.#####...##.##
#...####.###...
##..#.###...#..
#...##.##...#.#
#.#.##.###.#.##
####.##.#..#.##
##.#.##.#..#.##`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  405,
		},
		{
			name:  "actual",
			input: input,
			want:  30518,
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
			want:  400,
		},
		{
			name:  "example2",
			input: example2,
			want:  1400,
		},
		{
			name:  "actual",
			input: input,
			want:  36735,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
