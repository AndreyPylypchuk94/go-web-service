package main

import (
	"fmt"
	"strings"
)

func main() {
	list := []int{1, 2, 3}

	list = NewStream(list).
		Map(func(el int) int { return el * el }).
		Filter(func(el int) bool { return el >= 5 }).
		ToSlice()

	listS := []string{"1", "2", "3"}

	listS = NewStream(listS).
		Map(func(el string) string { return el + el }).
		Filter(func(el string) bool { return strings.HasPrefix(el, "1") }).
		ToSlice()

	fmt.Println(listS)
}

type Stream[T any] struct {
	source []T
}

func NewStream[T any](source []T) *Stream[T] {
	return &Stream[T]{source: source}
}

func (s *Stream[T]) Map(f func(el T) T) *Stream[T] {
	result := make([]T, len(s.source))
	for i, x := range s.source {
		result[i] = f(x)
	}
	s.source = result
	return s
}

func (s *Stream[T]) Filter(f func(el T) bool) *Stream[T] {
	result := make([]T, 0)
	for _, x := range s.source {
		if f(x) {
			result = append(result, x)
		}
	}
	s.source = result
	return s
}

func (s *Stream[T]) ToSlice() []T {
	return s.source
}
