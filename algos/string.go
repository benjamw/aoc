package algos

import (
	"strings"
)

func Reverse(s string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
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

func ToSlice(s string) []string {
	return strings.Split(s, "\n")
}

func ToMultiSlice(s string) [][]string {
	slc := ToSlice(s)

	out := make([][]string, len(slc))
	for i, v := range slc {
		out[i] = strings.Fields(v)
	}

	return out
}

func ToColumnSlice(s string) [][]string {
	slc := ToMultiSlice(s)

	return Transpose(slc)
}
