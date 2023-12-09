package algos

import (
	"strings"
)

func Reverse(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

func TrimAll(slc []string) []string {
	for n := range slc {
		slc[n] = strings.TrimSpace(slc[n])
	}

	return slc
}
