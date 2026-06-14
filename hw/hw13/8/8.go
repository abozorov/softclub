package main

import "fmt"

type Storage[T any] struct {
	data map[int]T
}

func (s *Storage[T]) Add(id int, value T) {
	s.data[id] = value
}

func (s *Storage[T]) Get(id int) (T, bool) {
	v, ok := s.data[id]
	return v, ok
}

func (s *Storage[T]) Delete(id int) {
	delete(s.data, id)
}

func (s *Storage[T]) All() []T {
	a := make([]T, 0, len(s.data))

	for _, v := range s.data {
		a = append(a, v)
	}
	return a
}

func main() {
	st := new(Storage[string])
	st.data = make(map[int]string)

	st.Add(1, "Anush")
	st.Add(2, "Bozorov")
	st.Add(3, "Bob")
	fmt.Println(st.All())

	st.Delete(2)
	fmt.Println(st.All())
	fmt.Println(st.Get(1))
	fmt.Println(st.Get(2))
}
