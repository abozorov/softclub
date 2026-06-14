package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if !s.IsEmpty() {
		v := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return v, true
	}
	return *new(T), false
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func main() {
	st := new(Stack[string])
	st.Push("hjasjhhj")

	fmt.Println(st.Pop())
	fmt.Println(st.Pop())

	st.Push("yyyy")
	fmt.Println(st.IsEmpty())
}
