package main

import "fmt"

type Filter interface {
	Apply(nums []int) []int
}

type EvenFilter struct{}
type OddFilter struct{}
type PositiveFilter struct{}

func (e EvenFilter) Apply(a []int) []int {
	b := make([]int, 0, len(a))
	for _, v := range a {
		if v%2 == 0 {
			b = append(b, v)
		}
	}
	return b
}

func (e OddFilter) Apply(a []int) []int {
	b := make([]int, 0, len(a))
	for _, v := range a {
		if v%2 != 0 {
			b = append(b, v)
		}
	}
	return b
}

func (e PositiveFilter) Apply(a []int) []int {
	b := make([]int, 0, len(a))
	for _, v := range a {
		if v > 0 {
			b = append(b, v)
		}
	}
	return b
}

func main() {
	var f Filter

	f = EvenFilter{}
	fmt.Println(f.Apply([]int{1, -2, 3, 4, -5}))

	f = OddFilter{}
	fmt.Println(f.Apply([]int{1, -2, 3, 4, -5}))

	f = PositiveFilter{}
	fmt.Println(f.Apply([]int{1, -2, 3, 4, -5}))

}
