package main

import (
	"fmt"
	"strings"
)

type Operation interface {
    Apply(s string) string
}

type Upper struct{}
type AddPrefix struct{} 

func (u Upper) Apply(s string) string {
	return strings.ToUpper(s)
}
func (a AddPrefix) Apply(s string) string {
	return "Hello" + s
}

func ApplyAll(data map[string]Operation) {
	for k, v := range data {
		fmt.Printf("str:=%s, struct:=%v, Apply:=%s\n", k, v, v.Apply(k))
	}
}

func main() {
	data := map[string]Operation{
		"go": Upper{},
		"lang": Upper{},
		" world!": AddPrefix{},
	}
	ApplyAll(data)
}