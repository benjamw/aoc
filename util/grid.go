package util

import (
	"fmt"
	"strings"
)

func PrintGrid[T any](grid [][]T) {
	print(strings.Repeat("-", len(grid[0])+2) + "\n")
	for _, l := range grid {
		for _, v := range l {
			print(fmt.Sprintf("%v", v))
		}
		print("\n")
	}
	print(strings.Repeat("-", len(grid[0])+2) + "\n")
}

func PrintGridFunc[S ~[][]T, T any, V any](grid S, callable func(T) V) {
	print(strings.Repeat("-", len(grid[0])+2) + "\n")
	for _, l := range grid {
		for _, v := range l {
			print(fmt.Sprintf("%v", callable(v)))
		}
		print("\n")
	}
	print(strings.Repeat("-", len(grid[0])+2) + "\n")
}
