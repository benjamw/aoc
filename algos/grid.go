package algos

import (
	"strings"
)

type Direction int

const (
	North Direction = 1
	East  Direction = 2
	South Direction = 3
	West  Direction = 4
)

// MoveNorth moves all the movable elements in the grid northward as far as they will go
// before either hitting the edge of the grid, or an immovable element, or a movable element
// that has stopped because it hit another immovable element. i.e.- the movable elements will stack.
func MoveNorth[T comparable](grid [][]T, space T, immovable T, movable T) [][]T {
	// move across the grid from left to right
	for x := 0; x < len(grid[0]); x++ {
		// look along columns
		c := 0
		prevImmovable := -1 // start at the edge of the grid
		for y := 0; y < len(grid); y++ {
			switch grid[y][x] {
			case movable:
				grid[y][x] = space
				c++
			case immovable:
				// replace all removed elements at end of area, starting at the previously found immovable element
				grid = moveMovables(grid, x, prevImmovable, c, movable, North)

				prevImmovable = y
				c = 0
			default:
				continue
			}
		}

		// there may be movable elements left over, replace them
		grid = moveMovables(grid, x, prevImmovable, c, movable, North)
	}

	return grid
}

// MoveSouth moves all the movable elements in the grid southward as far as they will go
// before either hitting the edge of the grid, or an immovable element, or a movable element
// that has stopped because it hit another immovable element. i.e.- the movable elements will stack.
func MoveSouth[T comparable](grid [][]T, space T, immovable T, movable T) [][]T {
	// move across the grid from left to right
	for x := 0; x < len(grid[0]); x++ {
		// look along columns
		c := 0
		prevImmovable := len(grid) // start at the edge of the grid
		for y := len(grid) - 1; y >= 0; y-- {
			switch grid[y][x] {
			case movable:
				grid[y][x] = space
				c++
			case immovable:
				// replace all removed elements at end of area, starting at the previously found immovable element
				grid = moveMovables(grid, x, prevImmovable, c, movable, South)

				prevImmovable = y
				c = 0
			default:
				continue
			}
		}

		// there may be movable elements left over, replace them
		grid = moveMovables(grid, x, prevImmovable, c, movable, South)
	}

	return grid
}

// MoveWest moves all the movable elements in the grid westward as far as they will go
// before either hitting the edge of the grid, or an immovable element, or a movable element
// that has stopped because it hit another immovable element. i.e.- the movable elements will stack.
func MoveWest[T comparable](grid [][]T, space T, immovable T, movable T) [][]T {
	// move down the grid from top to bottom
	for y := 0; y < len(grid); y++ {
		// look along columns
		c := 0
		prevImmovable := -1 // start at the left (west) edge of the grid
		for x := 0; x < len(grid[0]); x++ {
			switch grid[y][x] {
			case movable:
				grid[y][x] = space
				c++
			case immovable:
				// replace all removed elements at end of area, starting at the previously found immovable element
				grid = moveMovables(grid, y, prevImmovable, c, movable, West)

				prevImmovable = x
				c = 0
			default:
				continue
			}
		}

		// there may be movable elements left over, replace them
		grid = moveMovables(grid, y, prevImmovable, c, movable, West)
	}

	return grid
}

// MoveEast moves all the movable elements in the grid eastward as far as they will go
// before either hitting the edge of the grid, or an immovable element, or a movable element
// that has stopped because it hit another immovable element. i.e.- the movable elements will stack.
func MoveEast[T comparable](grid [][]T, space T, immovable T, movable T) [][]T {
	// move down the grid from top to bottom
	for y := 0; y < len(grid); y++ {
		// look along columns
		c := 0
		prevImmovable := len(grid[0]) // start at the right (east) edge of the grid
		for x := len(grid[0]) - 1; x >= 0; x-- {
			switch grid[y][x] {
			case movable:
				grid[y][x] = space
				c++
			case immovable:
				// replace all removed elements at end of area, starting at the previously found immovable element
				grid = moveMovables(grid, y, prevImmovable, c, movable, East)

				prevImmovable = x
				c = 0
			default:
				continue
			}
		}

		// there may be movable elements left over, replace them
		grid = moveMovables(grid, y, prevImmovable, c, movable, East)
	}

	return grid
}

func moveMovables[T any](grid [][]T, rowCol int, prevImmovable int, count int, movable T, dir Direction) [][]T {
	switch dir {
	case North:
		for n := 1; n <= count; n++ {
			grid[prevImmovable+n][rowCol] = movable
		}
	case South:
		for n := 1; n <= count; n++ {
			grid[prevImmovable-n][rowCol] = movable
		}
	case East:
		for n := 1; n <= count; n++ {
			grid[rowCol][prevImmovable-n] = movable
		}
	case West:
		for n := 1; n <= count; n++ {
			grid[rowCol][prevImmovable+n] = movable
		}
	}

	return grid
}

func MakeGrid(input string) [][]string {
	ret := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		ret = append(ret, strings.Split(line, ""))
	}

	return ret
}

func Clone[T any](src [][]T) (dst [][]T) {
	dst = make([][]T, len(src))

	for y, row := range src {
		dst[y] = make([]T, len(row))
		for x, v := range row {
			dst[y][x] = v
		}
	}

	return
}
