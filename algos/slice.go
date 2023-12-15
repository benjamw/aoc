package algos

import (
	"reflect"
	"sort"
	"strconv"

	"github.com/benjamw/aoc/cast"
)

func ToIntSlice(in []string) []int {
	v := make([]int, 0, len(in))

	for _, value := range in {
		v = append(v, cast.ToInt(value))
	}

	return v
}

func ToInt64Slice(in []string) []int64 {
	v := make([]int64, 0, len(in))

	for _, value := range in {
		i, _ := strconv.ParseInt(value, 0, 64)
		v = append(v, i)
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

func RemoveIndex[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func RemoveElem[T any](s []T, elem T) []T {
	return RemoveIndex(s, Find(s, elem))
}

func Find[T any](s []T, v T) int {
	for i, t := range s {
		if reflect.DeepEqual(t, v) {
			return i
		}
	}

	return -1
}

func Transpose[T any](s [][]T) [][]T {
	xl := len(s[0])
	yl := len(s)
	ret := make([][]T, xl)
	for i := range ret {
		ret[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			ret[i][j] = s[j][i]
		}
	}
	return ret
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

func SumFunc[S ~[]T, T any](s S, callable func(t T) int) int {
	sum := 0
	for _, t := range s {
		sum += callable(t)
	}

	return sum
}

func ProdFunc[S ~[]T, T any](s S, callable func(t T) int) int {
	sum := 1
	for _, t := range s {
		sum *= callable(t)
	}

	return sum
}
