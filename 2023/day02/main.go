package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/benjamw/aoc/algos"
	"github.com/benjamw/aoc/cast"
)

//go:embed input.txt
var input string

var games map[int][]map[string]int
var maxRed int = 12
var maxGreen int = 13
var maxBlue int = 14

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	games = parseInput(input)

	badGames := make([]int, 0)
	goodGames := make([]int, 0)

	// loop through the games
GAMES:
	for n, grabs := range games {
		for _, grab := range grabs {
			for color, num := range grab {
				switch color {
				case "red":
					if num > maxRed {
						badGames = append(badGames, n)
						continue GAMES
					}
					break
				case "green":
					if num > maxGreen {
						badGames = append(badGames, n)
						continue GAMES
					}
					break
				case "blue":
					if num > maxBlue {
						badGames = append(badGames, n)
						continue GAMES
					}
					break
				}
			}
		}

		goodGames = append(goodGames, n)
	}

	return algos.Sum(goodGames) // 2406
}

func part2(input string) int {
	games = parseInput(input)

	var maxRed, maxGreen, maxBlue int

	powers := make([]int, 0)

	// loop through the games
	for _, grabs := range games {
		maxRed = 0
		maxGreen = 0
		maxBlue = 0
		for _, grab := range grabs {
			for color, num := range grab {
				switch color {
				case "red":
					if num > maxRed {
						maxRed = num
					}
					break
				case "green":
					if num > maxGreen {
						maxGreen = num
					}
					break
				case "blue":
					if num > maxBlue {
						maxBlue = num
					}
					break
				}
			}
		}

		powers = append(powers, maxRed*maxGreen*maxBlue)
	}

	return algos.Sum(powers) // 78375
}

func parseInput(input string) (ans map[int][]map[string]int) {
	ans = make(map[int][]map[string]int)

	for _, line := range strings.Split(input, "\n") {
		// split on :
		top := strings.Split(line, ":")
		top = algos.TrimAll(top)

		reTop := regexp.MustCompile("^Game\\s*(\\d+)$")
		matched := reTop.FindStringSubmatch(top[0])
		gameId := cast.ToInt(matched[1])

		// split on ;
		mid := strings.Split(top[1], ";")
		mid = algos.TrimAll(mid)

		grabs := make([]map[string]int, 0)
		for _, str := range mid {
			// split on ,
			grab := strings.Split(str, ",")
			grab = algos.TrimAll(grab)

			blockSet := make(map[string]int)
			for _, gr := range grab {
				// split on " "
				blocks := strings.Split(gr, " ")
				blocks = algos.TrimAll(blocks)

				blockSet[blocks[1]] = cast.ToInt(blocks[0])
			}

			grabs = append(grabs, blockSet)
		}

		ans[gameId] = grabs
	}

	return ans
}
