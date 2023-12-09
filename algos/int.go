package algos

func Sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func Product(arr []int) int {
	// empty products = 1, see 0! and x^0... https://en.wikipedia.org/wiki/Empty_product
	prod := 1
	for _, valueInt := range arr {
		prod *= valueInt
	}
	return prod
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int, integers ...int) int {
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
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func LCMs(s []int) int {
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

func GCDs(s []int) int {
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
