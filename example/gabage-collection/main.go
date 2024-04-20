package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v []byte
}

func main() {
	foos := make([]Foo, 1_000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two := keepFirstTwoEl(foos)

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func keepFirstTwoEl(foos []Foo) []Foo {
	// return foos[:2]
	res := make([]Foo, 2)
	copy(res, foos[:2])
	return res
}
func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
