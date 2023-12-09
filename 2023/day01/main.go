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
	nums := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		line = strings.Trim(line, "abcdefghijklmnopqrstuvwxyz")
		num := fmt.Sprintf("%s%s", string(line[0]), line[len(line)-1:])
		nums = append(nums, cast.ToInt(num))
	}

	return algos.Sum(nums) // 55090
}

func part2(input string) int {
	nums := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		line = replaceDigits(line)
		line = strings.Trim(line, "abcdefghijklmnopqrstuvwxyz")
		num := fmt.Sprintf("%s%s", string(line[0]), line[len(line)-1:])
		nums = append(nums, cast.ToInt(num))
	}

	return algos.Sum(nums) // 54845
}

func parseInput(input string) (ans string) {
	return input
}

func replaceDigits(l string) string {
	units := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	ref := regexp.MustCompile(fmt.Sprintf("^\\D*?(%s)", strings.Join(units, "|")))
	reb := regexp.MustCompile(fmt.Sprintf("^\\D*?(%s)", algos.Reverse(strings.Join(units, "|"))))

	l = string(ref.ReplaceAll([]byte(l), []byte("$1")))

	for idx, unit := range units {
		l = string(regexp.MustCompile(fmt.Sprintf("^%s", unit)).ReplaceAll([]byte(l), []byte(cast.ToString(idx))))
	}

	l = algos.Reverse(l)
	l = string(reb.ReplaceAll([]byte(l), []byte("$1")))
	l = algos.Reverse(l)

	for idx, unit := range units {
		l = string(regexp.MustCompile(fmt.Sprintf("%s$", unit)).ReplaceAll([]byte(l), []byte(cast.ToString(idx))))
	}

	return l
}
