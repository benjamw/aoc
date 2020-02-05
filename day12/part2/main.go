package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type moon struct {
	x, y, z, dx, dy, dz int
}
type oneDim struct {
	m, dm int
}

func main() {
	// stringSlice := readInputFile("../test2772.txt")
	stringSlice := readInputFile("../input.txt")
	// fmt.Println(stringSlice)

	// need to set x y z, dx, dy, dz of each of the starting moons
	sliceMoons := makeMoonSlice(stringSlice)
	// fmt.Println(sliceMoons)

	// manually make the three dimension slices, to be passed into stringify & iterate helper functions
	xDim := make([]oneDim, 0)
	yDim := make([]oneDim, 0)
	zDim := make([]oneDim, 0)

	for _, m := range sliceMoons {
		xDim = append(xDim, oneDim{m.x, m.dx})
		yDim = append(yDim, oneDim{m.y, m.dy})
		zDim = append(zDim, oneDim{m.z, m.dz})
	}

	// fmt.Println(xDim, yDim, zDim)

	// "stringify" them so they are easy to compare later
	initialX, initialY, initialZ := stringifyOneDim(xDim), stringifyOneDim(yDim), stringifyOneDim(zDim)

	// find the number of steps for each dimension to reach it's initial position & velocity
	xSteps, ySteps, zSteps := iterate(xDim, initialX), iterate(yDim, initialY), iterate(zDim, initialZ)
	// fmt.Println(xSteps, ySteps, zSteps)

	// print the final least common multiple of the three number of steps
	fmt.Println(lcm(xSteps, ySteps, zSteps))
}

// helper function to put the input file into a slice of strings (each elements is a line of the txt file)
func readInputFile(path string) []string {
	// var pixelString string
	resultSlice := make([]string, 0)
	absPath, _ := filepath.Abs(path)

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// pixelString = line
		resultSlice = append(resultSlice, line)
	}

	// return pixelString
	return resultSlice
}

// helper function to take the slice of strings and return a slice of moon structs
func makeMoonSlice(stringSlice []string) []moon {
	sliceMoons := make([]moon, 0)
	for _, str := range stringSlice {
		x := str[strings.Index(str, "x=")+2 : strings.Index(str, ",")]
		// this is gross
		y := str[strings.Index(str, "y=")+2 : strings.Index(str, "y=")+strings.Index(str[strings.Index(str, ",")+1:], ",")-1]
		z := str[strings.Index(str, "z=")+2 : len(str)-1]

		intx, _ := strconv.Atoi(x)
		inty, _ := strconv.Atoi(y)
		intz, _ := strconv.Atoi(z)
		sliceMoons = append(sliceMoons, moon{intx, inty, intz, 0, 0, 0})
	}
	return sliceMoons
}

// helper function that updates the velocity then coordinate of a slice of oneDim structs
func updateVelThenCoords(sliceOneDim []oneDim) {
	for start := 0; start < len(sliceOneDim); start++ {
		// update velocity
		// requires iterating through all of the moons again...
		for restIndex := start + 1; restIndex < len(sliceOneDim); restIndex++ {
			if sliceOneDim[start].m < sliceOneDim[restIndex].m {
				sliceOneDim[start].dm++
				sliceOneDim[restIndex].dm--
			} else if sliceOneDim[start].m > sliceOneDim[restIndex].m {
				sliceOneDim[start].dm--
				sliceOneDim[restIndex].dm++
			}
		}
	}

	// then update coordinates x y z
	for i2, e := range sliceOneDim {
		sliceOneDim[i2].m += e.dm
	}
}

// helper function that will stringify a slice of oneDims to compare its values to another
func stringifyOneDim(sliceOneDim []oneDim) (result string) {
	for _, m := range sliceOneDim {
		result += strconv.Itoa(m.m) + ","
		result += strconv.Itoa(m.dm) + ","
	}
	return result
}

// helper function that will return the number of steps until the initial state is reached
// uses string comparison and the stringifyOneDim helper function
func iterate(dims []oneDim, initialString string) int {
	for i := 0; ; i++ {
		updateVelThenCoords(dims)
		if stringifyOneDim(dims) == initialString {
			return i + 1
		}
	}
}

// helper function that returns the least common multiple of three integers
func lcm(x, y, z int) int {
	pFactX, pFactY, pFactZ := primeFactorization(x), primeFactorization(y), primeFactorization(z)
	fmt.Println(pFactX, pFactY, pFactZ)

	ans := 1

	// multiple by every value in every slice, but do not count duplicates
	for i, j, k := 0, 0, 0; i < len(pFactX) || j < len(pFactY) || k < len(pFactZ); {
		if i < len(pFactX) {
			ans *= pFactX[i]
			if pFactX[i] == pFactY[j] {
				j++
			}
			if pFactX[i] == pFactZ[k] {
				k++
			}
			i++
		} else if j < len(pFactY) {
			ans *= pFactY[j]
			if pFactY[j] == pFactZ[k] {
				k++
			}
			j++
		} else if k < len(pFactZ) {
			ans *= pFactZ[k]
			k++
		}
	}

	return ans
}

func primeFactorization(num int) []int {
	ans := make([]int, 0)
	for i := 2; num > 1; {
		if num%i == 0 {
			ans = append(ans, i)
			num /= i
		} else {
			i++
		}
	}
	return ans
}
