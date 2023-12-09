package structs

type Stack[T any] []T

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Push(v T) *[]T {
	n := make([]T, 0)
	n = append(n, v)
	n = append(n, *s...)
	return &n
}

func (s *Stack[T]) Pop() (t T, empty bool) {
	if s.IsEmpty() {
		return *new(T), true
	}

	v := (*s)[0]
	*s = (*s)[1:]
	return v, false
}

func (s *Stack[T]) Peek() T {
	return (*s)[0]
}
