package stack

import "fmt"

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) String() string {
	string_s := "["
	for _, e := range s.arr {
		string_s = fmt.Sprintf("%s%#v, ", string_s, e)
	}
	if string_s == "[" {
		return ""
	}
	string_s = string_s[:len(string_s)-2]
	string_s = string_s + "]"
	return string_s
}

func (s *Stack[T]) Push(e ...T) {
	s.arr = append(s.arr, e...)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.arr) < 1 {
		var noop T
		return noop, false
	}
	top := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return top, true
}

func (s *Stack[T]) Top() (T, bool) {
	if len(s.arr) < 1 {
		var noop T
		return noop, false
	}
	return s.arr[len(s.arr)-1], true
}
