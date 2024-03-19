package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello, OTUS!"), reverse.Int(24601))

	vals := map[string]float64{"a": 0.01, "b": 0.03, "c": 0.02}
	fmt.Println(SumIntsOrFlouts(vals))
}

type Number interface {
	int | int8 | int32 | int64 | float32 | float64
}

func SumIntsOrFlouts[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}
