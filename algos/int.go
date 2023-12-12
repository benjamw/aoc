package algos

func Sum[T ~uint | ~int |
	~int8 | ~int16 | ~int32 | ~int64 |
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64,
](arr []T) T {
	sum := *new(T)
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Product[T ~uint | ~int |
	~int8 | ~int16 | ~int32 | ~int64 |
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64,
](arr []T) T {
	// empty products = 1, see 0! and x^0... https://en.wikipedia.org/wiki/Empty_product
	var prod T
	for _, v := range arr {
		prod *= v
	}
	return prod
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD[T ~uint | ~int |
	~int8 | ~int16 | ~int32 | ~int64 |
	~uint8 | ~uint16 | ~uint32 | ~uint64,
](a, b T, integers ...T) T {
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

// find Least Common Multiple (LCM) via GCD
func LCM[T ~uint | ~int |
	~int8 | ~int16 | ~int32 | ~int64 |
	~uint8 | ~uint16 | ~uint32 | ~uint64,
](a, b T, integers ...T) T {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func LCMs[T ~uint | ~int |
	~int8 | ~int16 | ~int32 | ~int64 |
	~uint8 | ~uint16 | ~uint32 | ~uint64,
](s []T) T {
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

func GCDs[T ~uint | ~int |
	~int8 | ~int16 | ~int32 | ~int64 |
	~uint8 | ~uint16 | ~uint32 | ~uint64,
](s []T) T {
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
