package algos

import (
	"sort"

	"github.com/benjamw/aoc/cast"
)

func ToIntSlice(in []string) []int {
	v := make([]int, 0, len(in))

	for _, value := range in {
		v = append(v, cast.ToInt(value))
	}

	return v
}

func ToStringSlice(in []int) []string {
	v := make([]string, 0, len(in))

	for _, value := range in {
		v = append(v, cast.ToString(value))
	}

	return v
}

func RemoveIndex(s []interface{}, index int) []interface{} {
	ret := make([]interface{}, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// SimpleGeneric will find the intersection between two slices using comparison
// Simple has complexity: O(n^2)
func SimpleGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

// SortedGeneric will find the intersection between two slices using sorted slices
// Sorted has complexity: O(n * log(n)), a needs to be sorted
func SortedGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool {
			return b[i] == v
		})
		if idx < len(b) && b[idx] == v {
			set = append(set, v)
		}
	}

	return set
}

// HashGeneric will find the intersection between two slices using hashes
// Hash has complexity: O(n * x) where x is a factor of hash function efficiency (between 1 and 2)
func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
