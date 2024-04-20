package main

import (
	"fmt"
)

func main() {
	i := 0
	if true {
		i := 1 // shadowing`i` in this block
		fmt.Println(i)
	}
	fmt.Println(i)

	// f := float64(math.MaxFloat64*1.1) + float64(1.0)
	// println(f)

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	println(len(m))

	n := map[string]int{
		"a": 1,
		"b": 2,
	}
	println(len(n))

}
