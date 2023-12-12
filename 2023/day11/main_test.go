package main

import (
	"testing"
)

var example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  374,
		},
		{
			name:  "actual",
			input: input,
			want:  9274989,
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
		exp   int
		want  int
	}{
		{
			name:  "example",
			input: example,
			exp:   10,
			want:  1030,
		},
		{
			name:  "example",
			input: example,
			exp:   100,
			want:  8410,
		},
		{
			name:  "actual",
			input: input,
			exp:   1e6,
			want:  357134560737,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.exp); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
