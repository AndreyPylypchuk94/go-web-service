package main

import (
	"fmt"
	"reflect"
)

type ServiceA struct {
}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (a *ServiceA) A() {
	fmt.Println("a")
}

type ServiceB struct {
	a *ServiceA
}

func NewServiceB() *ServiceB {
	return &ServiceB{
		a: Get(&ServiceA{}).(*ServiceA),
	}
}

func (b *ServiceB) B() {
	b.a.A()
}

func main() {
	a := NewServiceA()
	Add(a)

	b := NewServiceB()
	Add(b)

	b.B()

	fmt.Println(beans)
}

var beans = make([]interface{}, 0)

func Add(bean interface{}) {
	beans = append(beans, bean)
}

func Get(t interface{}) interface{} {
	for _, el := range beans {
		if reflect.TypeOf(el).String() == reflect.TypeOf(t).String() {
			return el
		}
	}
	panic("not found bean")
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
