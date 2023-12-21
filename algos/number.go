package algos

import (
	"cmp"
)

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Int | Float
}

func Sum[T Number](arr []T) T {
	sum := *new(T) // 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumFunc[S ~[]T, T Number](s S, callable func(t T) T) T {
	sum := *new(T) // 0
	for _, t := range s {
		sum += callable(t)
	}

	return sum
}

func Product[T Number](arr []T) T {
	// empty products = 1, see 0! and x^0... https://en.wikipedia.org/wiki/Empty_product
	prod := *new(T) + 1
	for _, v := range arr {
		prod *= v
	}
	return prod
}

func ProdFunc[S ~[]T, T Number](s S, callable func(t T) T) T {
	prod := *new(T) + 1
	for _, t := range s {
		prod *= callable(t)
	}

	return prod
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD[T Int](a, b T, integers ...T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	result := a

	for i := 0; i < len(integers); i++ {
		result = GCD(result, integers[i])
	}

	return result
}

// GCDs GCD for slices
func GCDs[T Int](s []T) T {
	switch len(s) {
	case 0:
		return 0
	case 1:
		return s[0]
	case 2:
		return GCD(s[0], s[1])
	default:
		return GCD(s[0], s[1], s[2:]...)
	}
}

// LCM find the Least Common Multiple (LCM) via GCD
func LCM[T Int](a, b T, integers ...T) T {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// LCMs LCM for slices
func LCMs[T Int](s []T) T {
	switch len(s) {
	case 0:
		return 0
	case 1:
		return s[0]
	case 2:
		return LCM(s[0], s[1])
	default:
		return LCM(s[0], s[1], s[2:]...)
	}
}

// Min get the minimum value in the slice
func Min[T cmp.Ordered](s []T) T {
	m := s[0]
	for _, t := range s {
		if t < m {
			m = t
		}
	}

	return m
}

// Max get the maximum value in the slice
func Max[T cmp.Ordered](s []T) T {
	m := s[0]
	for _, t := range s {
		if t > m {
			m = t
		}
	}

	return m
}
